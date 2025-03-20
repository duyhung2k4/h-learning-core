package courseregisterhandle

import (
	constant "app/internal/constants"
	"app/internal/entity"
	middlewareapp "app/internal/middleware"
	routerconfig "app/internal/router_config"
	query "app/pkg/query/basic"
	rawquery "app/pkg/query/raw"

	"github.com/gin-gonic/gin"
)

type courseRegisterHandle struct {
	query    query.QueryService[entity.CourseRegister]
	rawQuery rawquery.QueryRawService[entity.CourseRegister]
}

type CourseRegisterHandle interface {
	Create(ctx *gin.Context)
	Detail(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

func NewHandle() CourseRegisterHandle {
	return &courseRegisterHandle{
		query:    query.Register[entity.CourseRegister](),
		rawQuery: rawquery.Register[entity.CourseRegister](),
	}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.POST_HTTP,
		Endpoint: "course-register/create",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.Create,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.GET_HTTP,
		Endpoint: "course-register/all",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.GetAll,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.GET_HTTP,
		Endpoint: "course-register/detail",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.Detail,
	})
}
