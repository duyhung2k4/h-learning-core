package coursehandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"

	"github.com/gin-gonic/gin"
)

func (h *courseHandle) GetAllCourse(ctx *gin.Context) {
	courses, err := h.service.QueryCourse.Find(requestdata.QueryReq[entity.Course]{
		Order: "id ASC",
	})

	if err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger("get-all-course", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, courses)
}
