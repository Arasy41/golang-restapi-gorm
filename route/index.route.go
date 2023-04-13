package route

import (
	"github.com/Artzy1401/clone-cineplex-backend-4/config"
	"github.com/Artzy1401/clone-cineplex-backend-4/controller"
	"github.com/Artzy1401/clone-cineplex-backend-4/middleware"
	"github.com/gofiber/fiber/v2"
)


func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", controller.LoginController)
	//r.Post("/register", controller.RegisterController)

	// FILM & COMMENT
	r.Get("/films", controller.FilmControllerGetAll)
	r.Get("/film/:id", controller.FilmControllerGetById)
	r.Post("/film", controller.FilmControllerCreate)
	r.Put("/film/:id", controller.FilmControllerUpdate)
	r.Delete("/film/:id", controller.FilmControllerDelete)
	r.Put("/film/:id/like", controller.FilmControllerLikeUpdate)
	r.Get("/comments", middleware.Auth, controller.CommentControllerGetComments)
	r.Post("/comment",middleware.Auth, controller.CommentControllerCreate)
	r.Delete("/comment/:id",middleware.Auth, controller.CommentControllerDelete)

	// USER
	r.Get("/user", middleware.Auth, controller.UserControllerGetAll)
	r.Get("/user/:id", controller.UserControllerGetById)
	r.Post("/user", controller.UserControllerCreate)
	r.Put("/user/:id", controller.UserControllerUpdate)
	r.Put("/user/:id/update-email", controller.UserControllerUpdateEmail)
	r.Delete("/user/:id", controller.UserControllerDelete)

	// THEATER
	r.Get("/theater", controller.TheaterControllerGetAll)
	r.Get("/theater/:id", controller.TheaterControllerGetById)
	r.Get("/films/theaterlist", controller.FilmHandlerGetByTheaterId)
	r.Get("/theaterdetails", controller.TheaterControllerGetDetails)
	r.Post("/theater", controller.TheaterControllerCreate)
	r.Post("/theaterlist", controller.TheaterControllerCreateTheaterList)
	r.Put("/theater/:id", controller.TheaterControllerUpdate)
	r.Delete("/theater/:id", controller.TheaterControllerDelete)

}