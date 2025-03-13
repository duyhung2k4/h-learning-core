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

func (h *quizzHandle) GetQuizzById(ctx *gin.Context) {
	quizzIdString := ctx.Query("id")
	if quizzIdString == "" {
		logapp.Logger(constant.TITLE_GET_PAYLOAD, apperrors.ErrorQuizzIdInvalid.Error(), constant.ERROR_LOG)
		httpresponse.BadRequest(ctx, apperrors.ErrorQuizzIdInvalid)
		return
	}

	quizzId, err := strconv.Atoi(quizzIdString)
	if err != nil {
		logapp.Logger("convert-quizz-id", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	res, err := h.service.QueryQuizz.First(requestdata.QueryReq[entity.Quizz]{
		Condition: "id = ?",
		Args:      []interface{}{quizzId},
	})
	if err != nil {
		logapp.Logger("get-quizz-by-id", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	httpresponse.Success(ctx, res)
}
