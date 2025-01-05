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

func (h *quizzHandle) UpdateQuizz(ctx *gin.Context) {
	var updateQuizzPayload requestdata.UpdateQuizzRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&updateQuizzPayload); err != nil {
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		httpresponse.BadRequest(ctx, err)
		return
	}

	_, err := h.service.GrpcClientQuizz.UpdateQuizz(ctx, &servicegrpc.UpdateQuizzRequest{
		Id: updateQuizzPayload.Id,
		Payload: &servicegrpc.UpdateQuizzPayload{
			Ask:        updateQuizzPayload.Ask,
			Time:       int32(updateQuizzPayload.Time),
			ResultType: enumgrpc.ResultType(enumgrpc.ResultType_value[string(updateQuizzPayload.ResultType)]),
			Result:     updateQuizzPayload.Result,
			Option:     updateQuizzPayload.Option,
		},
	})
	if err != nil {
		logapp.Logger("update-quizz-grpc", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	httpresponse.Success(ctx, nil)
}
