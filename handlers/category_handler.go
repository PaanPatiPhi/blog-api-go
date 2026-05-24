package handlers

import (
	"blog-api-go/models"
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

func GetCategoryByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	category, err := repositories.GetCategoryByID(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to get category: " + err.Error()})
	}
	return c.JSON(category)
}

func CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	if err := repositories.CreateCategory(category); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create category: " + err.Error()})
	}
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	if err := repositories.UpdateCategory(uint(id), category); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update category: " + err.Error()})
	}
	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := repositories.DeleteCategory(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete category: " + err.Error()})
	}
	return c.SendStatus(204)
}
