package api

import (
	"log"
	"net/http"
	"strings"

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
}

func parseFileName(line string) string {
	token := "filename="
	i := strings.Index(line, token)
	start := i + len(token)

	value := string(line[start:])
	return strings.Trim(value, `"`)
}
