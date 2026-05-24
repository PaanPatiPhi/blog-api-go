package main

import (
	"log"

	"blog-api-go/database"
	"blog-api-go/handlers"
	"blog-api-go/routes"

	// "blog-api-go/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	// database.DB.AutoMigrate(&models.User{})

	app := fiber.New()
	api := app.Group("/")

app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:5173",
    AllowHeaders: "Origin, Content-Type, Accept, Authorization", // ← เพิ่ม Authorization
    AllowMethods: "GET, POST, PUT, DELETE",
}))

	routes.SetupPostRoutes(app)
	routes.SetupCategoryRoutes(app)
	routes.SetupCommentRoutes(app)

	api.Get("/users", handlers.GetUsers)



	log.Fatal(app.Listen(":4002"))
}
