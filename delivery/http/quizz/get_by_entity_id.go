package quizzhandle

import (
	"app/internal/apperrors"
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *quizzHandle) GetQuizzByEntityId(ctx *gin.Context) {
	entityIdString := ctx.Query("id")
	entityTypeString := ctx.Query("type")

	if entityIdString == "" {
		httpresponse.BadRequest(ctx, apperrors.ErrorQuizzEntityIdInvalid)
		return
	}
	if entityTypeString == "" {
		httpresponse.BadRequest(ctx, apperrors.ErrorQuizzEntityTypeInvalid)
		return
	}

	entityId, err := strconv.Atoi(entityIdString)
	if err != nil {
		logapp.Logger("convert-entity-id", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	res, err := h.service.QueryQuizz.Find(requestdata.QueryReq[entity.Quizz]{
		Condition: "entity_id = ?",
		Args:      []interface{}{entityId},
	})
	if err != nil {
		logapp.Logger("get-quizz", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	httpresponse.Success(ctx, res)
}
