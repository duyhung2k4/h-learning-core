package authhandle

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

func (h *authHandle) UpdateProfile(ctx *gin.Context) {
	var payload requestdata.UpdateProfileRequest

	// Parse request body
	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		httpresponse.BadRequest(ctx, err)
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		return
	}

	// Get profile ID from context
	profileId, exists := ctx.Get(string(constant.PROFILE_ID_KEY))
	if !exists {
		httpresponse.Unauthorized(ctx, errors.New("unauthorized"))
		logapp.Logger("get-profile-id", "profile ID not found in context", constant.ERROR_LOG)
		return
	}

	// Update profile
	updatedProfile, err := h.service.AuthService.UpdateProfile(ctx, profileId.(uint), payload)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("update-profile", err.Error(), constant.ERROR_LOG)
		return
	}

	httpresponse.Success(ctx, updatedProfile)
}
