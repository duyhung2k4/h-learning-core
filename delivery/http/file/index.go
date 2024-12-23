package filehandle

import (
	constant "app/internal/constants"
	routerconfig "app/internal/router_config"

	"github.com/gin-gonic/gin"
)

type fileHandle struct{}

type FileHandle interface {
	Thumnail(ctx *gin.Context)
}

func NewHandle() FileHandle {
	return &fileHandle{}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.GET_HTTP,
		Endpoint:   "file/thumnail_course/:filename",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.Thumnail,
	})
}
