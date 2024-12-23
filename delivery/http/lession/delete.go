package lessionhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *lessionHandle) Delete(ctx *gin.Context) {
	var payload requestdata.DeleteLessionReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	_, err := h.service.RawQueryLession.Query(requestdata.QueryRawReq[entity.Lession]{
		Args: []interface{}{
			time.Now(),
			payload.Id,
			profileId,
		},
		Sql: `
			UPDATE lessions
			SET
				deleted_at = ?
			FROM courses
			WHERE
				lessions.id = ?
				AND lessions.course_id = courses.id
				AND courses.create_id = ?
		`,
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("delete-lession", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, nil)
}
