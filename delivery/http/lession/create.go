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

func (h *lessionHandle) Create(ctx *gin.Context) {
	var payload requestdata.CreateLessionReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	lastLession, err := h.service.RawQueryLession.Query(requestdata.QueryRawReq[entity.Lession]{
		Sql: `
			SELECT l.* FROM lessions AS l
			JOIN chapters AS ct ON ct.id = l.chapter_id
			JOIN courses AS c ON c.id = l.course_id
			WHERE
				c.id = ?
				AND ct.id = ?
				AND c.create_id = ?
			ORDER BY l.order DESC
			LIMIT 1
		`,
		Args: []interface{}{
			payload.CourseId,
			payload.ChapterId,
			profileId,
		},
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("last-lession", err.Error(), constant.ERROR_LOG)
		return
	}

	lession, err := h.service.QueryLession.Create(entity.Lession{
		Name:        payload.Name,
		Description: payload.Description,
		CourseId:    payload.CourseId,
		ChapterId:   payload.ChapterId,
		Order:       lastLession.Order + 1,
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("create-lession", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, lession)
}
