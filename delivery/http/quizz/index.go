package quizzhandle

import (
	"app/cmd/core-service/service"
	constant "app/internal/constants"
	routerconfig "app/internal/router_config"

	"github.com/gin-gonic/gin"
)

type quizzHandle struct {
	service service.Service
}

type QuizzHandle interface {
	CreateQuizz(ctx *gin.Context)
	UpdateQuizz(ctx *gin.Context)
	DeleteQuizz(ctx *gin.Context)
	GetQuizzById(ctx *gin.Context)
	GetQuizzByEntityId(ctx *gin.Context)
}

func NewHandle() QuizzHandle {
	return &quizzHandle{
		service: service.Register(),
	}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.GET_HTTP,
		Endpoint:   "quizz/get-by-id",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.GetQuizzById,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.GET_HTTP,
		Endpoint:   "quizz/get-by-entity-id",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.GetQuizzByEntityId,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.POST_HTTP,
		Endpoint:   "quizz/create",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.CreateQuizz,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.PUT_HTTP,
		Endpoint:   "quizz/update",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.UpdateQuizz,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.DELETE_HTTP,
		Endpoint:   "quizz/delete",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.DeleteQuizz,
	})
}
