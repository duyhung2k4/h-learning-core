package authhandle

import (
	constant "app/internal/constants"
	responsedata "app/internal/dto/server"
	httpresponse "app/pkg/http_response"
	jwtapp "app/pkg/jwt"
	logapp "app/pkg/log"

	"github.com/gin-gonic/gin"
)

func (h *authHandle) RefreshToken(ctx *gin.Context) {
	profileId := ctx.GetUint(string(constant.PROFILE_ID_KEY))

	profile, err := h.service.AuthService.GetProfile(ctx, profileId)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("get-profile", err.Error(), constant.ERROR_LOG)
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
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
		Profile:      *profile,
	}

	httpresponse.Success(ctx, res)
}
