package quizzhandle

import (
	"app/generated/grpc/sharedgrpc"
	"app/internal/apperrors"
	constant "app/internal/constants"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	res, err := h.service.GrpcClientQuizz.GetById(ctx, &sharedgrpc.ID{
		Id: uint64(quizzId),
	})
	if err != nil {
		logapp.Logger("get-quizz-by-id", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	dataResponse := entity.Quizz{
		Model: gorm.Model{
			ID:        uint(res.Id),
			CreatedAt: time.Unix(res.CreatedAt, 0).Local(),
			UpdatedAt: time.Unix(res.UpdatedAt, 0).Local(),
		},
		Ask:        res.Ask,
		Time:       int(res.Time),
		ResultType: entity.RESULT_TYPE(res.ResultType.String()),
		Result:     res.Result,
		Option:     res.Option,
		EntityType: entity.ENTITY_TYPE(res.Option[res.EntityType]),
		EntityId:   uint(res.EntityId),
	}

	httpresponse.Success(ctx, dataResponse)
}
