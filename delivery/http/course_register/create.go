package courseregisterhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *courseRegisterHandle) Create(ctx *gin.Context) {
	var payload requestdata.CreateCourseRegisterReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		httpresponse.BadRequest(ctx, err)
		return
	}

	profileId := ctx.GetUint("profile_id")

	course, err := h.query.First(requestdata.QueryReq[entity.CourseRegister]{
		Condition: "course_id = ? AND profile_id = ?",
		Args:      []interface{}{payload.CourseId, profileId},
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		logapp.Logger("check-exist-course-register", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}
	if course != nil {
		errNil := errors.New("course exit")
		logapp.Logger("course-register-exist", errNil.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, errNil)
		return
	}

	result, err := h.query.Create(entity.CourseRegister{
		ProfileId: profileId,
		CourseId:  payload.CourseId,
	})
	if err != nil {
		logapp.Logger("create-course-register", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	httpresponse.Success(ctx, result)
}
