package authservice

import (
	"app/internal/entity"

	"github.com/gin-gonic/gin"
)

func (s *authService) GetProfile(ctx *gin.Context, profileId uint) (*entity.Profile, error) {
	var profile *entity.Profile

	if err := s.psql.
		Model(&entity.Profile{}).
		Preload("Role").
		Where("id = ?", profileId).
		First(&profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}
