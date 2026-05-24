package routes

import (
	"blog-api-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(app *fiber.App) {
	api := app.Group("/categories")
	api.Get("/", handlers.GetCategories)
}