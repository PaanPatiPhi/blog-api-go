package main

import (
	"log"

	"blog-api-go/database"
	"blog-api-go/handlers"
	"blog-api-go/routes"

	// "blog-api-go/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	// database.DB.AutoMigrate(&models.User{})

	app := fiber.New()
	api := app.Group("/api")

	api.Get("/users", handlers.GetUsers)
	routes.SetupPostRoutes(app)

	log.Fatal(app.Listen(":4003"))
}
