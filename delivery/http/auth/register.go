package authhandle

import (
	"app/internal/connection"
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	responsedata "app/internal/dto/server"
	httpresponse "app/pkg/http_response"
	jwtapp "app/pkg/jwt"
	logapp "app/pkg/log"
	"app/pkg/random"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *authHandle) Register(ctx *gin.Context) {
	var payload requestdata.RegisterRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	exist, err := h.service.AuthService.CheckExistAccount(ctx, payload.Email, payload.Phone)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("check-account", err.Error(), constant.ERROR_LOG)
		return
	}

	if *exist {
		httpresponse.InternalServerError(ctx, errors.New("email exist"))
		logapp.Logger("check-exist", "email exist", constant.ERROR_LOG)
		return
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("uuid", err.Error(), constant.ERROR_LOG)
		return
	}

	code := random.RandomCode(constant.LENGTH_CODE)

	err = h.service.AuthService.SaveInfoRegsiter(ctx, uuid.String(), code, payload)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("save-info-register", err.Error(), constant.ERROR_LOG)
		return
	}

	infoToken := map[string]interface{}{
		"uuid": uuid,
	}

	token, err := jwtapp.GenToken(infoToken, constant.ACCEPT_CODE_TOKEN_TIME)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("gen-token", err.Error(), constant.ERROR_LOG)
		return
	}

	h.emailJob.PushJob(connection.EmailJob_MessPayload{
		Email:   payload.Email,
		Content: code,
	})

	res := responsedata.RegisterResponse{
		Token: *token,
	}

	httpresponse.Success(ctx, res)
}
