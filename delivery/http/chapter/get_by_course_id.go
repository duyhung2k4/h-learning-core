package chapterhandle

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

func (h *chapterHandle) GetByCourseId(ctx *gin.Context) {
	courseIdString := ctx.Query("id")

	if courseIdString == "" {
		httpresponse.BadRequest(ctx, errors.New("id null"))
		logapp.Logger(constant.TITLE_GET_PAYLOAD, "id null", constant.ERROR_LOG)
		return
	}

	courseId, err := strconv.Atoi(courseIdString)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("convert-course-id", err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	chapters, err := h.service.QueryChapter.Find(requestdata.QueryReq[entity.Chapter]{
		Preload: map[string]*string{
			"Lessions": nil,
		},
		Joins: []string{
			"JOIN courses ON courses.id = chapters.course_id",
		},
		Condition: "courses.create_id = ? AND courses.id = ?",
		Args:      []interface{}{profileId, uint(courseId)},
		Order:     "chapters.order ASC",
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("get-course", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, chapters)
}
