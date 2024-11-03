package router

import (
	"app/config"
	"app/controller"
	middlewares "app/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func apiV1(router chi.Router) {
	authController := controller.NewAuthController()

	middlewares := middlewares.NewMiddlewares()

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
}
