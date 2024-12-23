package coursehandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"

	"github.com/gin-gonic/gin"
)

func (h *courseHandle) GetCourse(ctx *gin.Context) {
	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	courses, err := h.service.QueryCourse.Find(requestdata.QueryReq[entity.Course]{
		Condition: "create_id = ?",
		Args:      []interface{}{profileId},
		Order:     "id asc",
	})

	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("query", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, courses)
}
