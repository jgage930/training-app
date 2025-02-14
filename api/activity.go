package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/muktihari/fit/decoder"
)

type ActivityHandler struct {
	DB *sqlx.DB
}

func ActivityRouter(h *ActivityHandler, mux *http.ServeMux) {
	mux.HandleFunc("POST /activity/upload", h.uploadActivity)
	mux.HandleFunc("GET /activity/test", h.tester)
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

func (h *ActivityHandler) tester(w http.ResponseWriter, r *http.Request) {
	readFitFile("activity_data/2024-08-26-14-45-27.fit")
}

func readFitFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dec := decoder.New(f)

	fit, err := dec.Decode()
	if err != nil {
		panic(err)
	}

	log.Printf("FileHeader DataSize: %d\n", fit.FileHeader.DataSize)
	log.Printf("Messages count: %d\n", len(fit.Messages))
	// FileId is always the first message; 4 = activity

	// We want to track both activity and courses.
	// message = 4 or 5
	log.Printf("File Type: %v\n", fit.Messages[5]) //.FieldValueByNum(fieldnum.FileIdType).Any())
}
