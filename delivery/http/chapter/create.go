package chapterhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (h *chapterHandle) Create(ctx *gin.Context) {
	var payload requestdata.CreateChapterReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	lastChapter, err := h.service.RawQueryChapter.Query(requestdata.QueryRawReq[entity.Chapter]{
		Sql: `
			SELECT ct.* FROM
				chapters AS ct
			JOIN courses AS c ON c.id = ct.course_id
			WHERE 
				c.create_id = ? 
				AND c.id = ?
			ORDER BY ct.order DESC
			LIMIT 1
		`,
		Args: []interface{}{profileId, payload.CourseId},
	})

	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("get-last-chapter", err.Error(), constant.ERROR_LOG)
		return
	}

	chapter, err := h.service.QueryChapter.Create(entity.Chapter{
		Name:        payload.Name,
		Description: payload.Description,
		CourseId:    payload.CourseId,
		Order:       lastChapter.Order + 1,
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("get-chapter", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, chapter)
}
