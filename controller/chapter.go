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

type chapterController struct {
	queryChapterService    service.QueryService[model.Chapter]
	queryLessionService    service.QueryService[model.Lession]
	queryChapterRawService service.QueryRawService[model.Chapter]
	jwtUtils               utils.JwtUtils
}

type ChapterController interface {
	GetByCourseId(w http.ResponseWriter, r *http.Request)
	GetAllPublic(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (c *chapterController) GetByCourseId(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	courseIdString := params.Get("id")

	if courseIdString == "" {
		BadRequest(w, r, errors.New("courseId null"))
		return
	}

	courseId, err := strconv.Atoi(courseIdString)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	chapters, err := c.queryChapterService.Find(request.QueryReq[model.Chapter]{
		Preload: map[string]*string{
			"Lessions": nil,
		},
		Joins: []string{
			"JOIN courses ON courses.id = chapters.course_id",
		},
		Condition: "courses.create_id = ? AND courses.id = ?",
		Args:      []interface{}{profileId, uint(courseId)},
		Order:     "chapters.order ASC",
	})
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    chapters,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *chapterController) Create(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateChapterReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	lastChapter, err := c.queryChapterRawService.Query(request.QueryRawReq[model.Chapter]{
		Sql: `
			SELECT ct.* FROM
				chapters AS ct
			JOIN courses AS c ON c.id = ct.course_id
			WHERE 
				c.create_id = ? 
				AND c.id = ?
			ORDER BY ct.order DESC
			LIMIT 1
		`,
		Args: []interface{}{profileId, payload.CourseId},
	})

	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	chapter, err := c.queryChapterService.Create(model.Chapter{
		Name:        payload.Name,
		Description: payload.Description,
		CourseId:    payload.CourseId,
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

	newChapter, err := c.queryChapterRawService.Query(request.QueryRawReq[model.Chapter]{
		Sql: `
			UPDATE chapters
			SET
				name = ?,
				description = ?
			FROM courses
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

	_, err = c.queryChapterRawService.Query(request.QueryRawReq[model.Chapter]{
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

func (c *chapterController) GetAllPublic(w http.ResponseWriter, r *http.Request) {}

func NewChapterController() ChapterController {
	return &chapterController{
		queryChapterService:    service.NewQueryService[model.Chapter](),
		queryLessionService:    service.NewQueryService[model.Lession](),
		jwtUtils:               utils.NewJwtUtils(),
		queryChapterRawService: service.NewQueryRawService[model.Chapter](),
	}
}
