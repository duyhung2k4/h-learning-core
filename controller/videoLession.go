package controller

import "net/http"

type videoLessionController struct{}

type VideoLessionController interface {
	CreateVideoLession(w http.ResponseWriter, r *http.Request)
	DeleteVideoLession(w http.ResponseWriter, r *http.Request)
}

func (c *videoLessionController) CreateVideoLession(w http.ResponseWriter, r *http.Request) {}
func (c *videoLessionController) DeleteVideoLession(w http.ResponseWriter, r *http.Request) {}

func NewVideoLessionController() VideoLessionController {
	return &videoLessionController{}
}
