package authservice

import (
	constant "app/internal/constants"
	"app/internal/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *authService) CheckExistAccount(ctx *gin.Context, email string, phone string) (*bool, error) {
	var profile *entity.Profile

	if err := s.psql.
		Model(&entity.Profile{}).
		Where("email = ? AND phone = ?", email, phone).
		First(&profile).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if profile.ID != 0 {
		return &constant.TRUE, nil
	}

	return &constant.FALSE, nil
}
