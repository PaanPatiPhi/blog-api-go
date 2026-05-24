package routes

import (
	"blog-api-go/handlers"
	"blog-api-go/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(app *fiber.App) {
	api := app.Group("/categories")
	api.Get("/", handlers.GetCategories)
	api.Get("/:id", handlers.GetCategoryByID)
	api.Post("/", middlewares.AdminOnly, handlers.CreateCategory)
	api.Put("/:id", middlewares.AdminOnly, handlers.UpdateCategory)
	api.Delete("/:id", middlewares.AdminOnly, handlers.DeleteCategory)
}