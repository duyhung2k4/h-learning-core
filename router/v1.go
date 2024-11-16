package router

import (
	"app/config"
	"app/controller"
	middlewares "app/middlewares"
	"app/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func apiV1(router chi.Router) {

	middlewares := middlewares.NewMiddlewares()

	authController := controller.NewAuthController()

	courseQueryController := controller.NewQueryController[model.Course]()
	categoryQueryController := controller.NewQueryController[model.Category]()
	chapterQueryController := controller.NewQueryController[model.Chapter]()
	lessionQueryController := controller.NewQueryController[model.Lession]()
	documentLession := controller.NewQueryController[model.DocumentLession]()

	courseController := controller.NewCourseController()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]interface{}{
			"mess": "done",
		})
	})

	router.Route("/public", func(public chi.Router) {
	})

	router.Route("/auth", func(auth chi.Router) {
		auth.Post("/login", authController.Login)
		auth.Post("/register", authController.Register)
		auth.Post("/accept-code", authController.AcceptCopde)
	})

	router.Route("/protected", func(protected chi.Router) {
		protected.Use(jwtauth.Verifier(config.GetJWT()))
		protected.Use(middlewares.ValidateExpAccessToken())

		protected.Route("/auth", func(auth chi.Router) {
			auth.Post("/refresh-token", authController.RefreshToken)
		})

		protected.Route("/course", func(course chi.Router) {
			course.Post("/create", courseController.CreateCourse)
			course.Put("/update", courseController.UpdateCourse)
			course.Delete("/delete", courseController.DeleteCourse)
		})

		protected.Post("/chapter", chapterQueryController.Query)
		protected.Post("/lession", lessionQueryController.Query)
		protected.Post("/document-lession", documentLession.Query)
	})

	router.Route("/query", func(query chi.Router) {
		query.Post("/course", courseQueryController.Query)
		query.Post("/category", categoryQueryController.Query)
		query.Post("/chapter", chapterQueryController.Query)
		query.Post("/lession", lessionQueryController.Query)
		query.Post("/document-lession", documentLession.Query)
	})
}
