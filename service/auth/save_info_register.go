package authservice

import (
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *authService) SaveInfoRegsiter(ctx *gin.Context, uuid string, code string, infoRegister requestdata.RegisterRequest) error {
	data := SaveInfoRegisterPayload{
		InfoRegister: infoRegister,
		Code:         code,
	}

	jsonString, err := json.Marshal(data)

	if err != nil {
		return err
	}

	_, err = s.redis.Set(
		ctx,
		uuid,
		jsonString,
		time.Duration(constant.EXP_INFO_REGISTER)*time.Second,
	).Result()

	if err != nil {
		return err
	}

	return nil
}

type SaveInfoRegisterPayload struct {
	InfoRegister requestdata.RegisterRequest `json:"infoRegister"`
	Code         string                      `json:"code"`
}
