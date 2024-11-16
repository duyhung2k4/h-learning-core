package controller

import (
	"app/model"
	"app/service"
	"net/http"
)

type chapterController struct {
	queryService service.QueryService[model.Chapter]
}

type ChapterController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (c *chapterController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *chapterController) Update(w http.ResponseWriter, r *http.Request) {}

func (c *chapterController) Delete(w http.ResponseWriter, r *http.Request) {}

func NewChapterController() ChapterController {
	return &chapterController{
		queryService: service.NewQueryService[model.Chapter](),
	}
}
