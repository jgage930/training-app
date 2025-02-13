package api

import (
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
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
}
