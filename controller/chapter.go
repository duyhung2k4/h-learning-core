package controller

import (
	"app/dto/request"
	"app/model"
	"app/service"
	"app/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type chapterController struct {
	queryService    service.QueryService[model.Chapter]
	queryRawService service.QueryRawService[model.Chapter]
	jwtUtils        utils.JwtUtils
}

type ChapterController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (c *chapterController) Create(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateChapterReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	chapter, err := c.queryService.Create(model.Chapter{
		Name:        payload.Name,
		Description: payload.Description,
		CourseId:    payload.CourseId,
	})
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    chapter,
		Message: "OK",
		Error:   nil,
		Status:  200,
	}

	render.JSON(w, r, res)
}

func (c *chapterController) Update(w http.ResponseWriter, r *http.Request) {
	var payload request.UpdateChapterReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	newChapter, err := c.queryRawService.Query(request.QueryRawReq[model.Chapter]{
		Sql: `
			UPDATE chapters
			SET
				name = ?,
				description = ?
			FROM courses, profiles
			WHERE
				chapters.id = ?
				AND chapters.course_id = courses.id
  				AND courses.create_id = ?
			RETURNING chapters.*
		`,
		Data: []interface{}{payload.Name, payload.Description},
		Args: []interface{}{
			payload.Id,
			profileId,
		},
	})
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    newChapter,
		Message: "OK",
		Error:   nil,
		Status:  200,
	}

	render.JSON(w, r, res)
}

func (c *chapterController) Delete(w http.ResponseWriter, r *http.Request) {
	var payload request.DeleteChapterReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	_, err = c.queryRawService.Query(request.QueryRawReq[model.Chapter]{
		Args: []interface{}{
			time.Now(),
			payload.Id,
			profileId,
		},
		Sql: `
			UPDATE chapters
			SET
				deleted_at = ?
			FROM courses
			WHERE
				chapters.id = ?
				AND chapters.course_id = courses.id
				AND courses.create_id = ?
		`,
	})
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    nil,
		Message: "OK",
		Error:   nil,
		Status:  200,
	}

	render.JSON(w, r, res)
}

func NewChapterController() ChapterController {
	return &chapterController{
		queryService:    service.NewQueryService[model.Chapter](),
		jwtUtils:        utils.NewJwtUtils(),
		queryRawService: service.NewQueryRawService[model.Chapter](),
	}
}
