package authhandle

import (
	jobapp "app/cmd/core-service/job"
	"app/cmd/core-service/service"
	"app/internal/connection"
	constant "app/internal/constants"
	"app/internal/entity"
	middlewareapp "app/internal/middleware"
	routerconfig "app/internal/router_config"
	query "app/pkg/query/basic"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type authHandle struct {
	redis    *redis.Client
	service  service.Service
	query    query.QueryService[entity.Profile]
	emailJob jobapp.EmailJob
}

type AuthHandle interface {
	Login(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	Register(ctx *gin.Context)
	AcceptCopde(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

func NewAuthHandle() AuthHandle {
	return &authHandle{
		redis:    connection.GetRedisClient(),
		service:  service.Register(),
		emailJob: jobapp.NewEmailJob(),
		query:    query.Register[entity.Profile](),
	}
}

func Register(r *gin.Engine) {
	handle := NewAuthHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.POST_HTTP,
		Endpoint:   "auth/login",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.Login,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.POST_HTTP,
		Endpoint:   "auth/register",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.Register,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.POST_HTTP,
		Endpoint:   "auth/accept-code",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.AcceptCopde,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.POST_HTTP,
		Endpoint: "auth/refresh-token",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.RefreshToken,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.PUT_HTTP,
		Endpoint: "auth/update-profile",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.UpdateProfile,
	})
}
