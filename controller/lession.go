package controller

import (
	"app/dto/request"
	"app/model"
	"app/service"
	"app/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/render"
)

type lessionController struct {
	queryService    service.QueryService[model.Lession]
	queryRawService service.QueryRawService[model.Lession]
	jwtUtils        utils.JwtUtils
}

type LessionController interface {
	GetDetailLession(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (c *lessionController) GetDetailLession(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idString := params.Get("id")

	if idString == "" {
		InternalServerError(w, r, errors.New("id null"))
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	lession, err := c.queryService.First(request.QueryReq[model.Lession]{
		Preload: map[string]*string{
			"Chapter":      nil,
			"Course":       nil,
			"VideoLession": nil,
		},
		Joins: []string{
			"JOIN chapters AS ct ON ct.id = lessions.chapter_id",
			"JOIN courses AS c ON c.id = lessions.course_id",
		},
		Condition: `
			c.create_id = ?
			AND lessions.id = ?
		`,
		Args: []interface{}{
			profileId,
			id,
		},
	})

	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    lession,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *lessionController) Create(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateLessionReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	lastChapter, err := c.queryRawService.Query(request.QueryRawReq[model.Lession]{
		Sql: `
			SELECT l.* FROM lessions AS l
			JOIN chapters AS ct ON ct.id = l.chapter_id
			JOIN courses AS c ON c.id = l.course_id
			WHERE
				c.id = ?
				AND ct.id = ?
				AND c.create_id = ?
			ORDER BY l.order DESC
			LIMIT 1
		`,
		Args: []interface{}{
			payload.CourseId,
			payload.ChapterId,
			profileId,
		},
	})
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	chapter, err := c.queryService.Create(model.Lession{
		Name:        payload.Name,
		Description: payload.Description,
		CourseId:    payload.CourseId,
		ChapterId:   payload.ChapterId,
		Order:       lastChapter.Order + 1,
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

func (c *lessionController) Update(w http.ResponseWriter, r *http.Request) {
	var payload request.UpdateLessionReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	newChapter, err := c.queryRawService.Query(request.QueryRawReq[model.Lession]{
		Sql: `
			UPDATE lessions
			SET
				name = ?,
				description = ?
			FROM courses
			WHERE
				lessions.id = ?
				AND lessions.course_id = courses.id
  				AND courses.create_id = ?
			RETURNING lessions.*
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

func (c *lessionController) Delete(w http.ResponseWriter, r *http.Request) {
	var payload request.DeleteLessionReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	_, err = c.queryRawService.Query(request.QueryRawReq[model.Lession]{
		Args: []interface{}{
			time.Now(),
			payload.Id,
			profileId,
		},
		Sql: `
			UPDATE lessions
			SET
				deleted_at = ?
			FROM courses
			WHERE
				lessions.id = ?
				AND lessions.course_id = courses.id
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

func NewLessionController() LessionController {
	return &lessionController{
		queryService:    service.NewQueryService[model.Lession](),
		jwtUtils:        utils.NewJwtUtils(),
		queryRawService: service.NewQueryRawService[model.Lession](),
	}
}
