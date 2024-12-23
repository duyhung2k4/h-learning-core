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

func (h *courseHandle) GetDetailCourse(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		httpresponse.BadRequest(ctx, errors.New("id null"))
		logapp.Logger(constant.TITLE_GET_PAYLOAD, "id null", constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	courseId, err := strconv.Atoi(id)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("convert-id", err.Error(), constant.ERROR_LOG)
		return
	}

	course, err := h.service.QueryCourse.First(requestdata.QueryReq[entity.Course]{
		Condition: "id = ? AND create_id = ?",
		Args:      []interface{}{uint(courseId), profileId},
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("detail-course", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, course)
}
