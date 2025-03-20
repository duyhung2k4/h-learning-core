package lessionhandle

import (
	"app/cmd/core-service/service"
	constant "app/internal/constants"
	middlewareapp "app/internal/middleware"
	routerconfig "app/internal/router_config"

	"github.com/gin-gonic/gin"
)

type lessionHandle struct {
	service service.Service
}

type LessionHandle interface {
	GetDetailLession(ctx *gin.Context)
	GetDetailLessionPublic(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewHandle() LessionHandle {
	return &lessionHandle{
		service: service.Register(),
	}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.GET_HTTP,
		Endpoint: "lession/detail",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.GetDetailLession,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.GET_HTTP,
		Endpoint: "lession/detail-public",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.GetDetailLessionPublic,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.POST_HTTP,
		Endpoint: "lession/create",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.Create,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.PUT_HTTP,
		Endpoint: "lession/update",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.Update,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.DELETE_HTTP,
		Endpoint: "lession/delete",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.Delete,
	})
}
