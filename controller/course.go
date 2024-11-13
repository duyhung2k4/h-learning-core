package controller

import (
	"app/dto/request"
	"app/model"
	"app/service"
	"app/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type courseController struct {
	queryCourse service.QueryService[model.Course]
	jwtUtils    utils.JwtUtils
}

type CourseController interface {
	getProfileId(r *http.Request) (uint, error)

	CreateCourse(w http.ResponseWriter, r *http.Request)
	UpdateCourse(w http.ResponseWriter, r *http.Request)
	DeleteCourse(w http.ResponseWriter, r *http.Request)
}

func (c *courseController) getProfileId(r *http.Request) (uint, error) {
	token := c.jwtUtils.GetToken(r)
	mapInfo, err := c.jwtUtils.GetMapToken(token)

	if err != nil {
		return 0, err
	}

	profileId := uint(mapInfo["id"].(float64))

	return profileId, nil
}

func (c *courseController) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var payload request.QueryReq[model.Course]
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.getProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	codeCourse, err := uuid.NewV6()
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	payload.Data.CreateId = profileId
	payload.Data.Code = codeCourse.String()

	result, err := c.queryCourse.Create(payload.Data)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    result,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *courseController) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var payload request.QueryReq[model.Course]
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.getProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	payload.Data.CreateId = profileId
	payload.Condition = "id = ? AND create_id = ?"
	payload.Args = []interface{}{payload.Data.ID, profileId}

	result, err := c.queryCourse.Update(payload)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    result,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *courseController) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	var payload request.QueryReq[model.Course]
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.getProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	payload.Data.CreateId = profileId
	payload.Unscoped = false
	payload.Condition = "id = ? AND create_id = ?"
	payload.Args = []interface{}{payload.Data.ID, profileId}

	err = c.queryCourse.Delete(payload)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    nil,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func NewCourseController() CourseController {
	return &courseController{
		queryCourse: service.NewQueryService[model.Course](),
		jwtUtils:    utils.NewJwtUtils(),
	}
}
