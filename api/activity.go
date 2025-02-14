package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/profile/filedef"
)

type ActivityHandler struct {
	DB *sqlx.DB
}

type Activity struct {
	Id       int       `db:"id" json:"id"`
	FilePath string    `db:"file_path" json:"file_path"`
	Messages []Message `db:"messages" json:"messges"`
	// Need to parse some more data not sure what yet.
}

type Message struct {
	Id        int     `db:"id" json:"id"`
	Distance  float64 `db:"distance" json:"distance"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Longitude float64 `db:"longitude" json:"longitude"`
	Speed     float64 `db:"speed" json:"speed"`
	HeartRate uint8   `db:"heart_rate" json:"heart_rate"`
}

func ActivityRouter(h *ActivityHandler, mux *http.ServeMux) {
	mux.HandleFunc("GET /activity", h.listActivities)
	mux.HandleFunc("POST /activity/upload", h.uploadActivity)
}

func (h *ActivityHandler) listActivities(w http.ResponseWriter, r *http.Request) {
	activities := []Activity{}
	query := `
		SELECT * FROM activities;
	`
	h.DB.Select(&activities, query)
	Respond(activities, w)
}

func (h *ActivityHandler) uploadActivity(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(100 << 20)

	file, handler, err := r.FormFile("fileName")
	if err != nil {
		log.Printf("Error Retrieving File: \n%s", err)
		return
	}
	defer file.Close()

	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)

	path := "activity_data/" + handler.Filename
	localFile, err := os.Create(path)
	if err != nil {
		log.Printf("Error Creating file: \n%s", err)
		return
	}
	defer localFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error Reading file from request: \n%s", err)
		return
	}

	localFile.Write(fileBytes)
	localFile.Sync()

	// Parse fit file
	activity := readFitFile(path)
	// add to db
	tx := h.DB.MustBegin()
	result := tx.MustExec(
		"INSERT INTO activities (file_path) VALUES ($1)",
		activity.FilePath,
	)

	activity_id, err := result.LastInsertId()
	if err != nil {
		log.Println("Err could not get id.")
	}

	for _, message := range activity.Messages {
		tx.MustExec(
			`
			INSERT INTO activity_messages 
				(activity_id, distance, latitude, longitude, speed, heart_rate) 
			VALUES ($1, $2, $3, $4, $5, $6)
			`,
			activity_id,
			message.Distance,
			message.Latitude,
			message.Longitude,
			message.Speed,
			message.HeartRate,
		)
	}

	tx.Commit()

	log.Printf("Inserted %v", activity_id)
}

func readFitFile(path string) Activity {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// The listener will receive every message from the decoder
	// as soon as it is decoded and transform it into an filedef.File.
	lis := filedef.NewListener()
	defer lis.Close() // release channel used by listener

	dec := decoder.New(f,
		// Add activity listener to the decoder:
		decoder.WithMesgListener(lis),
		// Direct the decoder to only broadcast
		// the messages without retaining them:
		decoder.WithBroadcastOnly(),
	)

	_, err = dec.Decode()
	if err != nil {
		panic(err)
	}

	file, ok := lis.File().(*filedef.Activity)
	if !ok {
		log.Printf("%T is not an Activity File\n", lis.File())
	}

	var messages []Message
	for _, record := range file.Records {
		message := Message{
			Id:        0,
			Distance:  record.DistanceScaled(),
			Latitude:  record.PositionLatDegrees(),
			Longitude: record.PositionLongDegrees(),
			Speed:     record.SpeedScaled(),
			HeartRate: record.HeartRate,
		}

		messages = append(messages, message)
	}

	return Activity{
		Id:       0,
		FilePath: path,
		Messages: messages,
	}
}
