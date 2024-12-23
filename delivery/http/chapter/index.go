package chapterhandle

import (
	"app/cmd/core-service/service"
	constant "app/internal/constants"
	middlewareapp "app/internal/middleware"
	routerconfig "app/internal/router_config"

	"github.com/gin-gonic/gin"
)

type chapterHandle struct {
	service service.Service
}

type ChapterHandle interface {
	GetByCourseId(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func Newhandle() ChapterHandle {
	return &chapterHandle{
		service: service.Register(),
	}
}

func Register(r *gin.Engine) {
	handle := Newhandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.GET_HTTP,
		Endpoint: "chapter/get-by-course",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.GetByCourseId,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.POST_HTTP,
		Endpoint: "chapter/create",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.Create,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.PUT_HTTP,
		Endpoint: "chapter/update",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.Update,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.DELETE_HTTP,
		Endpoint: "chapter/delete",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.Delete,
	})
}
