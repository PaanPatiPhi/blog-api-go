package routes

import (
	"blog-api-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(app *fiber.App) {
    api := app.Group("/api")
    posts := api.Group("/posts")
    posts.Get("/published", handlers.GetPosts)
	posts.Get("/:id", handlers.GetPostByID)
}