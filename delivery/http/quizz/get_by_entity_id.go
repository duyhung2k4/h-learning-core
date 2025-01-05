package quizzhandle

import (
	"app/generated/proto/enumgrpc"
	"app/generated/proto/servicegrpc"
	"app/internal/apperrors"
	constant "app/internal/constants"
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

	h.service.GrpcClientQuizz.GetListByEntityId(ctx, &servicegrpc.GetListQuizzRequest{
		EntityId:   uint64(entityId),
		EntityType: enumgrpc.EntityType(enumgrpc.EntityType_value[string(entityIdString)]),
	})
}
