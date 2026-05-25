package routes

import (
	"blog-api-go/handlers"
	"blog-api-go/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupLikeRoutes(app *fiber.App) {
	api := app.Group("/likes")

	api.Get("/posts/:id/likes/count", handlers.GetLikesCount)
	api.Get("/posts/:id/likes/user/:userId", middlewares.AuthRequired, handlers.GetUserLikeStatus)
	api.Post("/posts/:id/likes", middlewares.AuthRequired, handlers.ToggleLike)	
}
