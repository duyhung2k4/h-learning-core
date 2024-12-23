package filehandle

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (c *fileHandle) Thumnail(ctx *gin.Context) {
	filename := ctx.Param("filename")
	filepath := fmt.Sprintf("cmd/core-service/data/file/thumnail_course/%s", filename)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	ctx.File(filepath)
}
