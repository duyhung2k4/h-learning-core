package coursehandle

import (
	"app/cmd/core-service/service"
	constant "app/internal/constants"
	middlewareapp "app/internal/middleware"
	routerconfig "app/internal/router_config"

	"github.com/gin-gonic/gin"
)

type courseHandle struct {
	service service.Service
}

type CourseHandle interface {
	GetCourse(ctx *gin.Context)
	GetDetailCourse(ctx *gin.Context)
	CreateCourse(ctx *gin.Context)
	UpdateCourse(ctx *gin.Context)
	ChangeActive(ctx *gin.Context)

	GetAllCourse(ctx *gin.Context)
	GetDetailCoursePublic(ctx *gin.Context)
}

func NewHandle() CourseHandle {
	return &courseHandle{
		service: service.Register(),
	}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.POST_HTTP,
		Endpoint: "course/create",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.CreateCourse,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.PUT_HTTP,
		Endpoint: "course/update",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.UpdateCourse,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.PUT_HTTP,
		Endpoint: "course/change-active",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.ChangeActive,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.GET_HTTP,
		Endpoint: "course/get-all",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.GetCourse,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.GET_HTTP,
		Endpoint: "course/detail",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.GetDetailCourse,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.GET_HTTP,
		Endpoint: "course/all-public",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.GetAllCourse,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.GET_HTTP,
		Endpoint: "course/detail-public",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.GetDetailCoursePublic,
	})
}
