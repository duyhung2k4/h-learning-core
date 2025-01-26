package documentlessionhandle

import (
	constant "app/internal/constants"
	"app/internal/entity"
	middlewareapp "app/internal/middleware"
	routerconfig "app/internal/router_config"
	query "app/pkg/query/basic"

	"github.com/gin-gonic/gin"
)

type documentLessionHandle struct {
	query query.QueryService[entity.DocumentLession]
}

type DocumentLessionHandle interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewHandle() DocumentLessionHandle {
	return &documentLessionHandle{
		query: query.Register[entity.DocumentLession](),
	}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method: constant.POST_HTTP,
		Handle: handle.Create,
		Middleware: []gin.HandlerFunc{
			middlewareapp.GetProfileId,
			middlewareapp.ValidateToken,
		},
		Endpoint: "document-lession/create",
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method: constant.PUT_HTTP,
		Handle: handle.Update,
		Middleware: []gin.HandlerFunc{
			middlewareapp.GetProfileId,
			middlewareapp.ValidateToken,
		},
		Endpoint: "document-lession/update",
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method: constant.DELETE_HTTP,
		Handle: handle.Delete,
		Middleware: []gin.HandlerFunc{
			middlewareapp.GetProfileId,
			middlewareapp.ValidateToken,
		},
		Endpoint: "document-lession/delete",
	})
}
