package authservice

import (
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	"context"
)

func (s *authService) UpdateProfile(ctx context.Context, profileId uint, payload requestdata.UpdateProfileRequest) (*entity.Profile, error) {
	var profile *entity.Profile
	err := s.psql.Model(entity.Profile{}).Where("id = ?", profileId).First(&profile).Error
	if err != nil {
		return nil, err
	}

	// Cập nhật các trường được phép
	if payload.FirstName != "" {
		profile.FirstName = payload.FirstName
	}
	if payload.LastName != "" {
		profile.LastName = payload.LastName
	}
	if payload.Phone != "" {
		profile.Phone = payload.Phone
	}

	if err = s.psql.Model(entity.Profile{}).Where("id = ?", profileId).Updates(&profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}
