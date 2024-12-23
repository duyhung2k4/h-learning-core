package lessionhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (h *lessionHandle) Update(ctx *gin.Context) {
	var payload requestdata.UpdateLessionReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	newLession, err := h.service.RawQueryLession.Query(requestdata.QueryRawReq[entity.Lession]{
		Sql: `
			UPDATE lessions
			SET
				name = ?,
				description = ?
			FROM courses
			WHERE
				lessions.id = ?
				AND lessions.course_id = courses.id
  				AND courses.create_id = ?
			RETURNING lessions.*
		`,
		Data: []interface{}{payload.Name, payload.Description},
		Args: []interface{}{
			payload.Id,
			profileId,
		},
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("update-lession", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, newLession)
}
