package authservice

import (
	"app/internal/connection"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type authService struct {
	psql     *gorm.DB
	redis    *redis.Client
	rabbitmq *amqp091.Connection
}

type AuthService interface {
	CheckExistAccount(ctx *gin.Context, email string, phone string) (*bool, error)
	GetProfile(ctx *gin.Context, profileId uint) (*entity.Profile, error)
	SaveInfoRegsiter(ctx *gin.Context, uuid string, code string, infoRegister requestdata.RegisterRequest) error
	CreateProfile(ctx *gin.Context, uuid string) error
	CompareProfile(ctx *gin.Context, payload requestdata.LoginRequest) (*entity.Profile, error)
	UpdateProfile(ctx context.Context, profileId uint, payload requestdata.UpdateProfileRequest) (*entity.Profile, error)
}

func Register() AuthService {
	return &authService{
		psql:     connection.GetPsql(),
		redis:    connection.GetRedisClient(),
		rabbitmq: connection.GetRabbitmq(),
	}
}
