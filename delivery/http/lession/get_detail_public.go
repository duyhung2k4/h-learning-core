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

func (h *lessionHandle) GetDetailLessionPublic(ctx *gin.Context) {
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

	lession, err := h.service.QueryLession.First(requestdata.QueryReq[entity.Lession]{
		Preload: map[string]*string{
			"Chapter":      nil,
			"Course":       nil,
			"VideoLession": nil,
		},
		Condition: "id = ?",
		Args: []interface{}{
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
