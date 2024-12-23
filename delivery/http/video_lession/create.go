package videolessionhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *videoLessionHandle) CreateVideoLession(ctx *gin.Context) {
	var payload requestdata.CreateVideoLessionReq
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
		Condition: "video_lessions.lession_id = ? AND c.create_id = ?",
		Args: []interface{}{
			payload.LessionId,
			profileId,
		},
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("check-exist", err.Error(), constant.ERROR_LOG)
		return
	}

	if videoLession != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("check-exist", "video has been initialized", constant.ERROR_LOG)
		return
	}

	uuidVideo, err := uuid.NewUUID()
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("uuid-video", err.Error(), constant.ERROR_LOG)
		return
	}

	result, err := h.service.QueryVideoLession.Create(entity.VideoLession{
		LessionId: payload.LessionId,
		Code:      uuidVideo.String(),
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("create-video-lession", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, result)
}
