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

func (h *quizzHandle) DeleteQuizz(ctx *gin.Context) {
	var deleteQuizzRequest requestdata.DeleteQuizzRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&deleteQuizzRequest); err != nil {
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		httpresponse.BadRequest(ctx, err)
		return
	}

	err := h.service.QueryQuizz.Delete(requestdata.QueryReq[entity.Quizz]{
		Condition: "id = ?",
		Args:      []interface{}{deleteQuizzRequest.Id},
	})
	if err != nil {
		logapp.Logger("delete-quizz-grpc", err.Error(), constant.ERROR_LOG)
		httpresponse.BadRequest(ctx, err)
		return
	}

	httpresponse.Success(ctx, nil)
}
