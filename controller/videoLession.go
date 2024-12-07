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

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type videoLessionController struct {
	jwtUtils        utils.JwtUtils
	queryService    service.QueryService[model.VideoLession]
	queryRawService service.QueryRawService[model.VideoLession]
}

type VideoLessionController interface {
	GetDetailVideoLession(w http.ResponseWriter, r *http.Request)
	CreateVideoLession(w http.ResponseWriter, r *http.Request)
	DeleteVideoLession(w http.ResponseWriter, r *http.Request)
	CheckVideoUpload(w http.ResponseWriter, r *http.Request)
}

func (c *videoLessionController) GetDetailVideoLession(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idString := params.Get("id")

	if idString == "" {
		BadRequest(w, r, errors.New("id null"))
		return
	}

	lessionId, err := strconv.Atoi(idString)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	videoLession, err := c.queryService.First(request.QueryReq[model.VideoLession]{
		Joins: []string{
			"JOIN lessions AS l ON l.id = video_lessions.lession_id",
			"JOIN courses AS c ON c.id = l.course_id",
		},
		Condition: "video_lessions.lession_id = ? AND c.create_id = ?",
		Args: []interface{}{
			lessionId,
			profileId,
		},
	})

	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    videoLession,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *videoLessionController) CreateVideoLession(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateVideoLessionReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	videoLession, err := c.queryService.First(request.QueryReq[model.VideoLession]{
		Joins: []string{
			"JOIN lessions AS l ON l.id = video_lessions.lession_id",
			"JOIN courses AS c ON c.id = l.course_id",
		},
		Condition: "video_lessions.lession_id = ? AND c.create_id = ?",
		Args: []interface{}{
			payload.LessionId,
			profileId,
		},
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		InternalServerError(w, r, err)
		return
	}

	if videoLession != nil {
		InternalServerError(w, r, errors.New("video has been initialized"))
		return
	}

	uuidVideo, err := uuid.NewV6()
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	result, err := c.queryService.Create(model.VideoLession{
		LessionId: payload.LessionId,
		Code:      uuidVideo.String(),
	})
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

func (c *videoLessionController) DeleteVideoLession(w http.ResponseWriter, r *http.Request) {}

func (c *videoLessionController) CheckVideoUpload(w http.ResponseWriter, r *http.Request) {
	var payload request.CheckVideoUploadReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profileId, err := c.jwtUtils.GetProfileId(r)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	videoLession, err := c.queryService.First(request.QueryReq[model.VideoLession]{
		Joins: []string{
			"JOIN lessions AS l ON l.id = video_lessions.lession_id",
			"JOIN courses AS c ON c.id = lessions.course_id",
		},
		Condition: "video_lessions.id = ? AND c.create_id = ?",
		Args:      []interface{}{payload.VideoLessionId, profileId},
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		InternalServerError(w, r, err)
		return
	}

	if videoLession.ID != 0 {
		InternalServerError(w, r, errors.New("video uploaded"))
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

func NewVideoLessionController() VideoLessionController {
	return &videoLessionController{
		jwtUtils:        utils.NewJwtUtils(),
		queryService:    service.NewQueryService[model.VideoLession](),
		queryRawService: service.NewQueryRawService[model.VideoLession](),
	}
}
