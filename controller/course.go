package controller

import (
	"app/constant"
	"app/dto/request"
	"app/model"
	"app/service"
	"app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type courseController struct {
	queryCourse service.QueryService[model.Course]
	jwtUtils    utils.JwtUtils
	fileUtils   utils.FileUtils
}

type CourseController interface {
	GetCourse(w http.ResponseWriter, r *http.Request)
	GetDetailCourse(w http.ResponseWriter, r *http.Request)
	CreateCourse(w http.ResponseWriter, r *http.Request)
	UpdateCourse(w http.ResponseWriter, r *http.Request)
	ChangeActive(w http.ResponseWriter, r *http.Request)

	GetAllCourse(w http.ResponseWriter, r *http.Request)
	GetDetailCoursePublic(w http.ResponseWriter, r *http.Request)
}

func (c *courseController) GetCourse(w http.ResponseWriter, r *http.Request) {
	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	courses, err := c.queryCourse.Find(request.QueryReq[model.Course]{
		Condition: "create_id = ?",
		Args:      []interface{}{profileId},
		Order:     "id asc",
	})

	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    courses,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *courseController) GetDetailCourse(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")

	if id == "" {
		BadRequest(w, r, errors.New("error id"))
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	courseId, err := strconv.Atoi(id)
	if err != nil {
		BadRequest(w, r, err)
		return
	}

	course, err := c.queryCourse.First(request.QueryReq[model.Course]{
		Condition: "id = ? AND create_id = ?",
		Args:      []interface{}{uint(courseId), profileId},
	})
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    course,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
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
			MultiLogin:  &payload.MultiLogin,
			Value:       payload.Value,
			Introduce:   payload.Introduce,
			Thumnail:    fmt.Sprintf("%s%s", uuidThumnail.String(), ext),
			Active:      &constant.TRUE,
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

	oldCourse, err := c.queryCourse.First(request.QueryReq[model.Course]{
		Condition: "id = ? AND create_id = ?",
		Args:      []interface{}{payload.Id, profileId},
	})
	if err != nil {
		InternalServerError(w, r, err)
		return
	}
	oldThumnail := fmt.Sprintf("file/thumnail_course/%s", oldCourse.Thumnail)

	newCourse := request.QueryReq[model.Course]{
		Data: model.Course{
			Name:        *payload.Name,
			Description: *payload.Description,
			MultiLogin:  payload.MultiLogin,
			Value:       *payload.Value,
			Thumnail:    fmt.Sprintf("%s%s", uuidThumnail.String(), ext),
			Introduce:   *payload.Introduce,
		},
		Condition: "id = ? AND create_id = ?",
		Args:      []interface{}{payload.Id, profileId},
	}

	result, err := c.queryCourse.Update(newCourse)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	err = os.RemoveAll(oldThumnail)
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
			Active: &payload.Active,
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

func (c *courseController) GetAllCourse(w http.ResponseWriter, r *http.Request) {
	courses, err := c.queryCourse.Find(request.QueryReq[model.Course]{
		Order: "id ASC",
	})

	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    courses,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *courseController) GetDetailCoursePublic(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")

	if id == "" {
		BadRequest(w, r, errors.New("error id"))
		return
	}

	courseId, err := strconv.Atoi(id)
	if err != nil {
		BadRequest(w, r, err)
		return
	}

	course, err := c.queryCourse.First(request.QueryReq[model.Course]{
		Condition: "id = ?",
		Args: []interface{}{
			uint(courseId),
		},
	})
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    course,
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
