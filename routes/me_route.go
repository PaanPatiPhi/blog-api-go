package routes

import (
	"blog-api-go/handlers"
	"blog-api-go/middlewares"

	"github.com/gofiber/fiber/v2"
)	

func MeRoute(app *fiber.App) {
	api := app.Group("/profile")
	api.Get("/", middlewares.AuthRequired, handlers.GetMeProfile)
	api.Put("/", middlewares.AuthRequired, handlers.UpdateUserProfile)
}