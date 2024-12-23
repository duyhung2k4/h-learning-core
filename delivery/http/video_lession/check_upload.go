package videolessionhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *videoLessionHandle) CheckVideoUpload(ctx *gin.Context) {
	var payload requestdata.CheckVideoUploadReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	videoLession, err := h.service.QueryVideoLession.First(requestdata.QueryReq[entity.VideoLession]{
		Joins: []string{
			"JOIN lessions AS l ON l.id = video_lessions.lession_id",
			"JOIN courses AS c ON c.id = l.course_id",
		},
		Condition: "video_lessions.id = ? AND c.create_id = ?",
		Args:      []interface{}{payload.VideoLessionId, profileId},
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("get-video-lession", err.Error(), constant.ERROR_LOG)
		return
	}

	if videoLession.Url360p != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("get-video-lession", "video uploaded", constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, nil)
}
