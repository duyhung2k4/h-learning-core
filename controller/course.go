package controller

import (
	"app/dto/request"
	"app/model"
	"app/service"
	"app/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type courseController struct {
	queryCourse service.QueryService[model.Course]
	jwtUtils    utils.JwtUtils
	fileUtils   utils.FileUtils
}

type CourseController interface {
	CreateCourse(w http.ResponseWriter, r *http.Request)
	UpdateCourse(w http.ResponseWriter, r *http.Request)
	ChangeActive(w http.ResponseWriter, r *http.Request)
}

func (c *courseController) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateCourseReq
	metadata := r.FormValue("metadata")
	err := json.Unmarshal([]byte(metadata), &payload)

	if err != nil {
		BadRequest(w, r, err)
		return
	}

	file, header, err := r.FormFile("thumnail")
	if err != nil {
		BadRequest(w, r, err)
		return
	}

	uuidThumnail, err := uuid.NewV6()
	if err != nil {
		InternalServerError(w, r, err)
		return
	}
	dirSave := "file/thumnail_course"

	_, ext, err := c.fileUtils.CreateFile(uuidThumnail.String(), dirSave, file, header)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	codeCourse, err := uuid.NewV6()
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	newCourse := request.QueryReq[model.Course]{
		Data: model.Course{
			CreateId:    profileId,
			Code:        codeCourse.String(),
			Name:        payload.Name,
			Description: payload.Description,
			MultiLogin:  payload.MultiLogin,
			Value:       payload.Value,
			Thumnail:    fmt.Sprintf("%s%s", uuidThumnail.String(), ext),
			Active:      true,
		},
	}

	result, err := c.queryCourse.Create(newCourse.Data)
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
	var payload request.UpdateCourseReq
	metadata := r.FormValue("metadata")
	err := json.Unmarshal([]byte(metadata), &payload)

	if err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	newCourse := request.QueryReq[model.Course]{
		Data: model.Course{
			Name:        *payload.Name,
			Description: *payload.Description,
			MultiLogin:  *payload.MultiLogin,
			Value:       *payload.Value,
		},
		Condition: "id = ? AND create_id = ?",
		Args:      []interface{}{payload.Id, profileId},
	}

	result, err := c.queryCourse.Update(newCourse)
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

func (c *courseController) ChangeActive(w http.ResponseWriter, r *http.Request) {
	var payload request.ChangeAvticeCourseReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	newCourseCourse := request.QueryReq[model.Course]{
		Data: model.Course{
			Active: payload.Active,
		},
		Condition: "id = ? AND create_id = ?",
		Args:      []interface{}{payload.Id, profileId},
	}

	result, err := c.queryCourse.Update(newCourseCourse)
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

func NewCourseController() CourseController {
	return &courseController{
		queryCourse: service.NewQueryService[model.Course](),
		jwtUtils:    utils.NewJwtUtils(),
		fileUtils:   utils.NewFileUtils(),
	}
}
