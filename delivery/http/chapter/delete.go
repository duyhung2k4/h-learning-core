package chapterhandle

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

func (h *chapterHandle) Delete(ctx *gin.Context) {
	var payload requestdata.UpdateChapterReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	_, err := h.service.RawQueryChapter.Query(requestdata.QueryRawReq[entity.Chapter]{
		Args: []interface{}{
			time.Now(),
			payload.Id,
			profileId,
		},
		Sql: `
			UPDATE chapters
			SET
				deleted_at = ?
			FROM courses
			WHERE
				chapters.id = ?
				AND chapters.course_id = courses.id
				AND courses.create_id = ?
		`,
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("dalete-chapter", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, nil)
}
