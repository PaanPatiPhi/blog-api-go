package handlers

import (
	"blog-api-go/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	categories, err := repositories.GetCategories()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to get categories: " + err.Error()})
	}
	return c.JSON(categories)
}
