package quizzhandle

import (
	"app/generated/proto/enumgrpc"
	"app/generated/proto/servicegrpc"
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
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

	res, err := h.service.GrpcClientQuizz.CreateQuizz(ctx, &servicegrpc.CreateQuizzRequest{
		Ask:        createQuizzRequest.Ask,
		Time:       int32(createQuizzRequest.Time),
		Option:     createQuizzRequest.Option,
		ResultType: enumgrpc.ResultType(enumgrpc.ResultType_value[string(createQuizzRequest.ResultType)]),
		Result:     createQuizzRequest.Result,
		EntityType: enumgrpc.EntityType(enumgrpc.EntityType_value[string(createQuizzRequest.EntityType)]),
		EntityId:   uint64(createQuizzRequest.EntityId),
	})

	if err != nil {
		logapp.Logger("create-quizz", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	httpresponse.Success(ctx, res)
}
