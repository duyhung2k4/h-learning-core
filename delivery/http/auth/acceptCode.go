package authhandle

import (
	authservice "app/cmd/core-service/service/auth"
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	httpresponse "app/pkg/http_response"
	jwtapp "app/pkg/jwt"
	logapp "app/pkg/log"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *authHandle) AcceptCopde(ctx *gin.Context) {
	var payload requestdata.AcceptCodeRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	token := jwtapp.GetToken(ctx)
	mapInfoToken, err := jwtapp.GetClaim(*token)
	if err != nil {
		httpresponse.Unauthorized(ctx, err)
		logapp.Logger("gen-token-accept-code", err.Error(), constant.ERROR_LOG)
		return
	}

	uuid := fmt.Sprint(mapInfoToken["uuid"])
	val, err := h.redis.Get(ctx, uuid).Result()
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("create-uuid", err.Error(), constant.ERROR_LOG)
		return
	}

	var infoRegister authservice.SaveInfoRegisterPayload
	err = json.Unmarshal([]byte(val), &infoRegister)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("map-json", err.Error(), constant.ERROR_LOG)
		return
	}

	if payload.Code != infoRegister.Code {
		httpresponse.InternalServerError(ctx, errors.New("code wrong"))
		logapp.Logger("compare-code", "code wrong", constant.ERROR_LOG)
		return
	}

	err = h.service.AuthService.CreateProfile(ctx, uuid)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("create-profile", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, nil)
}
