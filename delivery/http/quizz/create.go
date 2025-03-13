package quizzhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (h *quizzHandle) CreateQuizz(ctx *gin.Context) {
	var createQuizzRequest requestdata.CreateQuizzRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&createQuizzRequest); err != nil {
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		httpresponse.BadRequest(ctx, err)
		return
	}

	res, err := h.service.QueryQuizz.Create(entity.Quizz{
		Ask:        createQuizzRequest.Ask,
		ResultType: createQuizzRequest.ResultType,
		Result:     createQuizzRequest.Result,
		Option:     createQuizzRequest.Option,
		Time:       createQuizzRequest.Time,
		EntityType: createQuizzRequest.EntityType,
		EntityId:   createQuizzRequest.EntityId,
	})

	if err != nil {
		logapp.Logger("create-quizz", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	httpresponse.Success(ctx, res)
}
