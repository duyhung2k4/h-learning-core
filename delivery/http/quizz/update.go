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

func (h *quizzHandle) UpdateQuizz(ctx *gin.Context) {
	var updateQuizzPayload requestdata.UpdateQuizzRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&updateQuizzPayload); err != nil {
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		httpresponse.BadRequest(ctx, err)
		return
	}

	res, err := h.service.QueryQuizz.Update(requestdata.QueryReq[entity.Quizz]{
		Condition: "id = ?",
		Args:      []interface{}{updateQuizzPayload.Id},
		Data: entity.Quizz{
			Ask:        updateQuizzPayload.Ask,
			Time:       updateQuizzPayload.Time,
			ResultType: updateQuizzPayload.ResultType,
			Result:     updateQuizzPayload.Result,
			Option:     updateQuizzPayload.Option,
		},
	})

	if err != nil {
		logapp.Logger("update-quizz-grpc", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	httpresponse.Success(ctx, res)
}
