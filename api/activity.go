package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
)

type ActivityHandler struct {
	DB *sqlx.DB
}

type ActivityFile struct {
	FileName string
	Content  []byte
}

func ActivityRouter(h *ActivityHandler, mux *http.ServeMux) {
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
