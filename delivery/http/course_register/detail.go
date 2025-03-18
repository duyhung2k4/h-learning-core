package courseregisterhandle

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

func (h *courseRegisterHandle) Detail(ctx *gin.Context) {
	courseIdString := ctx.Query("id")
	courseId, err := strconv.Atoi(courseIdString)
	if err != nil {
		err := errors.New("courseId invalid")
		logapp.Logger("get-course-id", err.Error(), constant.ERROR_LOG)
		httpresponse.BadRequest(ctx, err)
		return
	}

	profileId := ctx.GetUint("profile_id")

	result, err := h.query.First(requestdata.QueryReq[entity.CourseRegister]{
		Condition: "profile_id = ? AND course_id = ?",
		Args:      []interface{}{profileId, courseId},
	})
	if err != nil {
		logapp.Logger("find-course-register", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	httpresponse.Success(ctx, result)
}
