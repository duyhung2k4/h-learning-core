package controller

import (
	"app/dto/request"
	"encoding/json"
	"net/http"
)

type quizzController struct{}

type QuizzController interface {
	CreateQuizz(w http.ResponseWriter, r *http.Request)
}

func (c *quizzController) CreateQuizz(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateQuizzRequest
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		BadRequest(w, r, err)
		return
	}
}

func NewQuizzController() QuizzController {
	return &quizzController{}
}
