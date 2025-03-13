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
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *courseHandle) UpdateCourse(ctx *gin.Context) {
	var payload requestdata.UpdateCourseReq
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
	dirSave := "data/thumbnail_course"
	_, ext, err := fileapp.CreateFile(uuidThumnail.String(), dirSave, file, header)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("dir-save", err.Error(), constant.ERROR_LOG)
		return
	}

	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	oldCourse, err := h.service.QueryCourse.First(requestdata.QueryReq[entity.Course]{
		Condition: "id = ? AND create_id = ?",
		Args:      []interface{}{payload.Id, profileId},
	})
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("get-old-course", err.Error(), constant.ERROR_LOG)
		return
	}
	oldThumnail := fmt.Sprintf("data/thumbnail_course/%s", oldCourse.Thumnail)

	newCourse := requestdata.QueryReq[entity.Course]{
		Data: entity.Course{
			Name:        *payload.Name,
			Description: *payload.Description,
			MultiLogin:  payload.MultiLogin,
			Value:       *payload.Value,
			Thumnail:    fmt.Sprintf("%s%s", uuidThumnail.String(), ext),
			Introduce:   *payload.Introduce,
		},
		Condition: "id = ? AND create_id = ?",
		Args:      []interface{}{payload.Id, profileId},
	}

	result, err := h.service.QueryCourse.Update(newCourse)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("update-course", err.Error(), constant.ERROR_LOG)
		return
	}

	err = os.RemoveAll(oldThumnail)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("remove-old-thumnail", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, result)
}
