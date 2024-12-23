package coursehandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (h *courseHandle) ChangeActive(ctx *gin.Context) {
	var payload requestdata.ChangeAvticeCourseReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	newCourseCourse := requestdata.QueryReq[entity.Course]{
		Data: entity.Course{
			Active: &payload.Active,
		},
		Condition: "id = ? AND create_id = ?",
		Args:      []interface{}{payload.Id, profileId},
	}

	result, err := h.service.QueryCourse.Update(newCourseCourse)
	if err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger("update-course", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, result)
}
