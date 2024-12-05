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
	chapterController := controller.NewChapterController()
	lessionController := controller.NewLessionController()
	documentLessionController := controller.NewDocumentLessionController()
	fileController := controller.NewFileController()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]interface{}{
			"mess": "done",
		})
	})

	router.Route("/public", func(public chi.Router) {
		public.Route("/file", func(file chi.Router) {
			file.Get("/thumnail_course/{filename}", fileController.Thumnail)
		})
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
			course.Get("/detail", courseController.GetDetailCourse)
			course.Get("/get-all", courseController.GetCourse)
			course.Post("/create", courseController.CreateCourse)
			course.Put("/update", courseController.UpdateCourse)
			course.Put("/change-active", courseController.ChangeActive)
		})

		protected.Route("/chapter", func(chapter chi.Router) {
			chapter.Get("/get-by-course", chapterController.GetByCourseId)
			chapter.Post("/create", chapterController.Create)
			chapter.Put("/update", chapterController.Update)
			chapter.Delete("/delete", chapterController.Delete)
		})

		protected.Route("/lession", func(lession chi.Router) {
			lession.Post("/create", lessionController.Create)
			lession.Put("/update", lessionController.Update)
			lession.Delete("/delete", lessionController.Delete)
		})

		protected.Route("/document-lession", func(documentLession chi.Router) {
			documentLession.Post("/create", documentLessionController.Create)
			documentLession.Put("/update", documentLessionController.Update)
			documentLession.Delete("/delete", documentLessionController.Delete)
		})
	})

	router.Route("/query", func(query chi.Router) {
		query.Post("/course", courseQueryController.Query)
		query.Post("/category", categoryQueryController.Query)
		query.Post("/chapter", chapterQueryController.Query)
		query.Post("/lession", lessionQueryController.Query)
		query.Post("/document-lession", documentLession.Query)
	})
}
