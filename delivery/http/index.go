package httpapp

import (
	authhandle "app/cmd/core-service/delivery/http/auth"
	chapterhandle "app/cmd/core-service/delivery/http/chapter"
	coursehandle "app/cmd/core-service/delivery/http/course"
	courseregisterhandle "app/cmd/core-service/delivery/http/course_register"
	documentlessionhandle "app/cmd/core-service/delivery/http/document_lession"
	filehandle "app/cmd/core-service/delivery/http/file"
	lessionhandle "app/cmd/core-service/delivery/http/lession"
	quizzhandle "app/cmd/core-service/delivery/http/quizz"
	videolessionhandle "app/cmd/core-service/delivery/http/video_lession"
	"app/internal/connection"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Register() http.Handler {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authhandle.Register(r)
	coursehandle.Register(r)
	chapterhandle.Register(r)
	lessionhandle.Register(r)
	videolessionhandle.Register(r)
	filehandle.Register(r)
	quizzhandle.Register(r)
	documentlessionhandle.Register(r)
	courseregisterhandle.Register(r)

	r.StaticFS("/file", http.Dir("data"))

	log.Printf(
		"Server h-learning-core starting success! URL: http://%s:%s",
		connection.GetConnect().CoreService.Host,
		connection.GetConnect().CoreService.Port,
	)

	return r
}
