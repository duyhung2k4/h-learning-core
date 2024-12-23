package videolessionhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *videoLessionHandle) GetDetailVideoLession(ctx *gin.Context) {
	lessionIdString := ctx.Query("id")

	if lessionIdString == "" {
		httpresponse.BadRequest(ctx, errors.New("id null"))
		logapp.Logger(constant.TITLE_GET_PAYLOAD, "id null", constant.ERROR_LOG)
		return
	}

	lessionId, err := strconv.Atoi(lessionIdString)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("convert-course-id", err.Error(), constant.ERROR_LOG)
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
			lessionId,
			profileId,
		},
	})

	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("video-lession-detail", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, videoLession)
}
