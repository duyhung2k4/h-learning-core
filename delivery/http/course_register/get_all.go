package courseregisterhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"

	"github.com/gin-gonic/gin"
)

func (h *courseRegisterHandle) GetAll(ctx *gin.Context) {
	profileId := ctx.GetUint("profile_id")

	result, err := h.query.Find(requestdata.QueryReq[entity.CourseRegister]{
		Condition: "profile_id = ?",
		Args:      []interface{}{profileId},
		Preload: map[string]*string{
			"Course": nil,
		},
	})
	if err != nil {
		logapp.Logger("get-all-course-register", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	httpresponse.Success(ctx, result)
}
