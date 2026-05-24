package routes

import (
	"blog-api-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupCommentRoutes(app *fiber.App) {
	api := app.Group("/comments")
	api.Post("/", handlers.CreateComment)
	api.Get("/post/:id", handlers.GetCommentsByPostID)
}
