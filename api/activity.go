package api

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/profile/filedef"
)

type ActivityHandler struct {
	DB *sqlx.DB
}

type Activity struct {
	Id   int       `db:"id" json:"id"`
	Date time.Time `db:"date" json:"date"`
	// Need to parse some more data not sure what yet.
}

func ActivityRouter(h *ActivityHandler, mux *http.ServeMux) {
	mux.HandleFunc("GET /activity", h.getActivity)
	mux.HandleFunc("POST /activity/upload", h.uploadActivity)
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
}

func (h *ActivityHandler) getActivity(w http.ResponseWriter, r *http.Request) {
	// file := r.PathValue("file")
	readFitFile("activity_data/2024-08-26-14-45-27.fit")
}

// func readFitFile(path string) {
// 	f, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
//
// 	dec := decoder.New(f)
//
// 	fit, err := dec.Decode()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	log.Printf("FileHeader DataSize: %d\n", fit.FileHeader.DataSize)
// 	log.Printf("Messages count: %d\n", len(fit.Messages))
// 	// FileId is always the first message; 4 = activity
//
// 	// We want to track both activity and courses.
// 	// message = 4 or 5
// 	log.Printf("File Type: %v\n", fit.Messages[5]) //.FieldValueByNum(fieldnum.FileIdType).Any())
//
// }

func readFitFile(path string) {
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
		return
	}

	log.Printf("File Type: %s\n", file.FileId.Type)
	log.Printf("Sessions count: %d\n", len(file.Sessions))
	log.Printf("Laps count: %d\n", len(file.Laps))
	log.Printf("Records count: %d\n", len(file.Records))
	i := 100
	log.Printf("\nSample value of record[%d]:\n", i)
	log.Printf("  Distance: %g m\n", file.Records[i].DistanceScaled())
	log.Printf("  Lat: %g degrees\n", file.Records[i].PositionLatDegrees())
	log.Printf("  Long: %g degrees\n", file.Records[i].PositionLongDegrees())
	log.Printf("  Speed: %g m/s\n", file.Records[i].SpeedScaled())
	log.Printf("  HeartRate: %d bpm\n", file.Records[i].HeartRate)
	log.Printf("  Cadence: %d rpm\n", file.Records[i].Cadence)
}
