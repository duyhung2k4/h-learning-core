package coursehandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	fileapp "app/pkg/file"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *courseHandle) CreateCourse(ctx *gin.Context) {
	var payload requestdata.CreateCourseReq
	metadata := ctx.Request.FormValue("metadata")
	err := json.Unmarshal([]byte(metadata), &payload)

	if err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	file, header, err := ctx.Request.FormFile("thumnail")
	if err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}
	uuidThumnail, err := uuid.NewUUID()
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("uuid-thumnail", err.Error(), constant.ERROR_LOG)
		return
	}
	dirSave := "cmd/core-service/data/file/thumnail_course"
	_, ext, err := fileapp.CreateFile(uuidThumnail.String(), dirSave, file, header)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("dir-save", err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	codeCourse, err := uuid.NewUUID()
	if err != nil {
		logapp.Logger("code-course", err.Error(), constant.ERROR_LOG)
		return
	}

	newCourse := requestdata.QueryReq[entity.Course]{
		Data: entity.Course{
			CreateId:    profileId,
			Code:        codeCourse.String(),
			Name:        payload.Name,
			Description: payload.Description,
			MultiLogin:  &payload.MultiLogin,
			Value:       payload.Value,
			Introduce:   payload.Introduce,
			Thumnail:    fmt.Sprintf("%s%s", uuidThumnail.String(), ext),
			Active:      &constant.TRUE,
		},
	}

	result, err := h.service.QueryCourse.Create(newCourse.Data)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("crete-course", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, result)
}
