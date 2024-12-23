package videolessionhandle

import (
	"app/cmd/core-service/service"
	constant "app/internal/constants"
	middlewareapp "app/internal/middleware"
	routerconfig "app/internal/router_config"

	"github.com/gin-gonic/gin"
)

type videoLessionHandle struct {
	service service.Service
}

type VideoLessionHandle interface {
	GetDetailVideoLession(ctx *gin.Context)
	CreateVideoLession(ctx *gin.Context)
	DeleteVideoLession(ctx *gin.Context)
	CheckVideoUpload(ctx *gin.Context)
}

func NewHandle() VideoLessionHandle {
	return &videoLessionHandle{
		service: service.Register(),
	}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.GET_HTTP,
		Endpoint: "video-lession/detail",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.GetDetailVideoLession,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.POST_HTTP,
		Endpoint: "video-lession/create",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.CreateVideoLession,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.DELETE_HTTP,
		Endpoint: "video-lession/delete",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.DeleteVideoLession,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.POST_HTTP,
		Endpoint: "video-lession/check-video-upload",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.CheckVideoUpload,
	})
}
