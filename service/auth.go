package service

import (
	"app/config"
	"app/constant"
	"app/dto/request"
	"app/model"
	"app/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type authService struct {
	psql      *gorm.DB
	redis     *redis.Client
	authUtils utils.AuthUtils
	jwtUtils  utils.JwtUtils
}

type AuthService interface {
	CheckExistAccount(email string, phone string) (*bool, error)
	CreateCode(length int) string
	GetProfile(profileId uint) (*model.Profile, error)
	SaveInfoRegsiter(uuid string, code string, infoRegister request.RegisterReq) error
	CreateProfile(uuid string) error
	CompareProfile(payload request.LoginRequest) (*model.Profile, error)
	CreateToken(data map[string]interface{}) (string, string, error)
}

func (s *authService) CheckExistAccount(email string, phone string) (*bool, error) {
	var profile *model.Profile

	if err := s.psql.
		Model(&model.Profile{}).
		Where("email = ? AND phone = ?", email, phone).
		First(&profile).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if profile.ID != 0 {
		return &constant.TRUE, nil
	}

	return &constant.FALSE, nil
}

func (s *authService) CreateCode(length int) string {
	b := ""
	for i := 0; i < length; i++ {
		b += fmt.Sprintf("%d", rand.Intn(9))
	}
	return b
}

func (s *authService) GetProfile(profileId uint) (*model.Profile, error) {
	var profile *model.Profile

	if err := s.psql.
		Model(&model.Profile{}).
		Preload("Role").
		Where("id = ?", profileId).
		First(&profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *authService) SaveInfoRegsiter(uuid string, code string, infoRegister request.RegisterReq) error {
	data := SaveInfoRegisterPayload{
		InfoRegister: infoRegister,
		Code:         code,
	}

	jsonString, err := json.Marshal(data)

	if err != nil {
		return err
	}

	_, err = s.redis.Set(
		context.Background(),
		uuid,
		jsonString,
		time.Duration(constant.EXP_INFO_REGISTER)*time.Second,
	).Result()

	if err != nil {
		return err
	}

	return nil
}

func (s *authService) CreateProfile(uuid string) error {
	var saveInfoRegister SaveInfoRegisterPayload
	jsonString, err := s.redis.Get(context.Background(), uuid).Result()

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(jsonString), &saveInfoRegister)
	if err != nil {
		return err
	}

	var role model.Role
	if err = s.psql.Model(&model.Role{}).Where("code = ?", model.USER).First(&role).Error; err != nil {
		return err
	}

	password, err := s.authUtils.HashPassword(saveInfoRegister.InfoRegister.Password)
	if err != nil {
		return err
	}

	var profile model.Profile = model.Profile{
		Email:    saveInfoRegister.InfoRegister.Email,
		Password: password,
		RoleId:   role.ID,
		Active:   true,
	}

	if err := s.psql.Model(&model.Profile{}).Create(&profile).Error; err != nil {
		return err
	}

	return nil
}

func (s *authService) CompareProfile(payload request.LoginRequest) (*model.Profile, error) {
	var profile *model.Profile

	err := s.psql.
		Model(&model.Profile{}).
		Preload("Role").
		Where("email = ?", payload.Username).
		First(&profile).Error
	if err != nil {
		return nil, err
	}

	isOk := s.authUtils.CheckPasswordHash(payload.Password, profile.Password)
	if !isOk {
		return nil, errors.New("password wrong")
	}

	return profile, nil
}

func (s *authService) CreateToken(data map[string]interface{}) (string, string, error) {
	var accessToken string
	var refreshToken string
	var err error

	uuidAccessToken, err := uuid.NewV6()
	if err != nil {
		return "", "", err
	}
	mapAccessToken := data
	mapAccessToken["uuid"] = uuidAccessToken
	mapAccessToken["exp"] = time.Now().Add(time.Second * time.Duration(constant.ACCESS_TOKEN_EXP))
	accessToken, err = s.jwtUtils.JwtEncode(mapAccessToken)
	if err != nil {
		return "", "", err
	}

	uuidRefreshToken, err := uuid.NewV6()
	if err != nil {
		return "", "", err
	}
	mapRefreshToken := data
	mapRefreshToken["uuid"] = uuidRefreshToken
	mapRefreshToken["exp"] = time.Now().Add(time.Second * time.Duration(constant.REFRESH_TOKEN_EXP))
	refreshToken, err = s.jwtUtils.JwtEncode(mapRefreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func NewAuthService() AuthService {
	return &authService{
		psql:      config.GetPsql(),
		redis:     config.GetRedisClient(),
		authUtils: utils.NewAuthUtils(),
		jwtUtils:  utils.NewJwtUtils(),
	}
}

type SaveInfoRegisterPayload struct {
	InfoRegister request.RegisterReq `json:"infoRegister"`
	Code         string              `json:"code"`
}
