package routes

import (
	"blog-api-go/handlers"
	"blog-api-go/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(app *fiber.App) {
	api := app.Group("/posts")

	api.Get("/published", handlers.GetPublishedPosts)
	api.Get("", middlewares.AdminOnly, handlers.GetAllPosts)
	api.Get("/:id", handlers.GetPostByID)
	api.Delete("/:id", middlewares.AdminOnly, handlers.DeletePostByID)
	api.Post("", middlewares.AdminOnly, handlers.CreatePost)
	api.Put("/:id", middlewares.AdminOnly, handlers.UpdatePost)
}