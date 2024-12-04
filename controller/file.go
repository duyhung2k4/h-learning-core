package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type fileController struct{}

type FileController interface {
	Thumnail(w http.ResponseWriter, r *http.Request)
}

func (c *fileController) Thumnail(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	filepath := fmt.Sprintf("file/thumnail_course/%s", filename)

	log.Println(filepath)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filepath)
}

func NewFileController() FileController {
	return &fileController{}
}
