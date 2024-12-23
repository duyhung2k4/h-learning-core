package lessionhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *lessionHandle) GetDetailLession(ctx *gin.Context) {
	lessionIdString := ctx.Query("id")

	if lessionIdString == "" {
		httpresponse.BadRequest(ctx, errors.New("id null"))
		logapp.Logger(constant.TITLE_GET_PAYLOAD, "id null", constant.ERROR_LOG)
		return
	}

	lessionId, err := strconv.Atoi(lessionIdString)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("convert-course-id", err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	lession, err := h.service.QueryLession.First(requestdata.QueryReq[entity.Lession]{
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
			lessionId,
		},
	})

	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("get-detail", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, lession)
}
