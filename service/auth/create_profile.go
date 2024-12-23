package authservice

import (
	constant "app/internal/constants"
	"app/internal/entity"
	"app/pkg/password"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (s *authService) CreateProfile(ctx *gin.Context, uuid string) error {
	var saveInfoRegister SaveInfoRegisterPayload
	jsonString, err := s.redis.Get(ctx, uuid).Result()

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(jsonString), &saveInfoRegister)
	if err != nil {
		return err
	}

	var role entity.Role
	if err = s.psql.Model(&entity.Role{}).Where("code = ?", entity.USER).First(&role).Error; err != nil {
		return err
	}

	password, err := password.HashPassword(saveInfoRegister.InfoRegister.Password)
	if err != nil {
		return err
	}

	var profile entity.Profile = entity.Profile{
		Email:    saveInfoRegister.InfoRegister.Email,
		Password: password,
		RoleId:   role.ID,
		Active:   &constant.TRUE,
	}

	if err := s.psql.Model(&entity.Profile{}).Create(&profile).Error; err != nil {
		return err
	}

	return nil
}
