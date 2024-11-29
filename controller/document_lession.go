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

type documentLessionController struct {
	queryService    service.QueryService[model.DocumentLession]
	queryRawService service.QueryRawService[model.DocumentLession]
	jwtUtils        utils.JwtUtils
}

type DocumentLessionController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (c *documentLessionController) Create(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateDocumentLessionReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	chapter, err := c.queryService.Create(model.DocumentLession{
		Content:   payload.Content,
		LessionId: payload.LessionId,
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

func (c *documentLessionController) Update(w http.ResponseWriter, r *http.Request) {
	var payload request.UpdateDocumentLessionReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	newChapter, err := c.queryRawService.Query(request.QueryRawReq[model.DocumentLession]{
		Sql: `
			UPDATE document_lessions
			SET
				content = ?
			FROM lessions, courses
			WHERE
				document_lessions.id = ?
				AND lessions.id = document_lessions.lession_id
				AND lessions.course_id = courses.id
  				AND courses.create_id = ?
			RETURNING document_lessions.*
		`,
		Data: []interface{}{payload.Content},
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

func (c *documentLessionController) Delete(w http.ResponseWriter, r *http.Request) {
	var payload request.DeleteDocumentLessionReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	_, err = c.queryRawService.Query(request.QueryRawReq[model.DocumentLession]{
		Args: []interface{}{
			time.Now(),
			payload.Id,
			profileId,
		},
		Sql: `
			UPDATE document_lessions
			SET
				deleted_at = ?
			FROM lessions, courses
			WHERE
				document_lessions.id = ?
				AND lessions.id = document_lessions.lession_id
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

func NewDocumentLessionController() DocumentLessionController {
	return &documentLessionController{
		queryService:    service.NewQueryService[model.DocumentLession](),
		jwtUtils:        utils.NewJwtUtils(),
		queryRawService: service.NewQueryRawService[model.DocumentLession](),
	}
}
