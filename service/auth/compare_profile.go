package authservice

import (
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	"app/pkg/password"
	"errors"

	"github.com/gin-gonic/gin"
)

func (s *authService) CompareProfile(ctx *gin.Context, payload requestdata.LoginRequest) (*entity.Profile, error) {
	var profile *entity.Profile

	err := s.psql.
		Model(&entity.Profile{}).
		Preload("Role").
		Where("email = ?", payload.Username).
		First(&profile).Error
	if err != nil {
		return nil, err
	}

	isOk := password.ComparePassword(payload.Password, profile.Password)
	if !isOk {
		return nil, errors.New("password wrong")
	}

	return profile, nil
}
