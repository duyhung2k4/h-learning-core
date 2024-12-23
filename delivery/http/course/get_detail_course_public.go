package coursehandle

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

func (h *courseHandle) GetDetailCoursePublic(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		httpresponse.BadRequest(ctx, errors.New("id null"))
		logapp.Logger(constant.TITLE_GET_PAYLOAD, "id null", constant.ERROR_LOG)
		return
	}

	courseId, err := strconv.Atoi(id)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("convert-course-id", err.Error(), constant.ERROR_LOG)
		return
	}

	course, err := h.service.QueryCourse.First(requestdata.QueryReq[entity.Course]{
		Condition: "id = ?",
		Args: []interface{}{
			uint(courseId),
		},
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("get-detail-course", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, course)
}
