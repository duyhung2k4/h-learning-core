package authhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	responsedata "app/internal/dto/server"
	httpresponse "app/pkg/http_response"
	jwtapp "app/pkg/jwt"
	logapp "app/pkg/log"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (h *authHandle) Login(ctx *gin.Context) {
	var payload requestdata.LoginRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	profile, err := h.service.AuthService.CompareProfile(ctx, payload)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("compare-password", err.Error(), constant.ERROR_LOG)
		return
	}

	mapInfo := map[string]interface{}{
		"email": profile.Email,
		"role":  profile.Role.Code,
		"id":    profile.ID,
	}

	accessToken, err := jwtapp.GenToken(mapInfo, constant.ACCESS_TOKEN_TIME)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("gen-access-token", err.Error(), constant.ERROR_LOG)
		return
	}
	refreshToken, err := jwtapp.GenToken(mapInfo, constant.ACCESS_TOKEN_TIME)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("gen-refresh-token", err.Error(), constant.ERROR_LOG)
		return
	}

	profile.Password = ""

	res := responsedata.LoginResponse{
		Profile:      *profile,
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}

	httpresponse.Success(ctx, res)
}
