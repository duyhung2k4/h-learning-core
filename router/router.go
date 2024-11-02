package router

import (
	//...

	"app/config"
	"app/controller"
	middlewares "app/middlewares"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func AppRouter() http.Handler {
	app := chi.NewRouter()

	// A good base middleware stack
	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	app.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	authController := controller.NewAuthController()

	middlewares := middlewares.NewMiddlewares()

	app.Route("/api/v1", func(router chi.Router) {

		router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, map[string]interface{}{
				"mess": "done",
			})
		})

		router.Route("/public", func(public chi.Router) {
			public.Post("/login", authController.Login)
		})

		router.Route("/auth", func(auth chi.Router) {
			auth.Post("/register", authController.Register)
		})

		router.Route("/protected", func(protected chi.Router) {
			protected.Use(jwtauth.Verifier(config.GetJWT()))
			protected.Use(jwtauth.Authenticator(config.GetJWT()))
			protected.Use(middlewares.ValidateExpAccessToken())

			protected.Post("/refresh-token", authController.RefreshToken)
		})

		router.Get("/file/save_auth/{filename}", func(w http.ResponseWriter, r *http.Request) {
			filename := chi.URLParam(r, "filename")
			imagePath := filepath.Join("file/save_auth", filename) // Thay đổi đường dẫn này

			// Kiểm tra nếu file tồn tại
			if _, err := os.Stat(imagePath); os.IsNotExist(err) {
				http.Error(w, "File not found", http.StatusNotFound)
				return
			}

			// Set header Content-Type để trình duyệt nhận diện đúng loại file
			w.Header().Set("Content-Type", "image/png") // Hoặc loại hình ảnh tương ứng
			http.ServeFile(w, r, imagePath)
		})
	})

	log.Printf(
		"Server art-pixel starting success! URL: http://%s:%s",
		config.GetAppHost(),
		config.GetAppPort(),
	)

	return app
}
