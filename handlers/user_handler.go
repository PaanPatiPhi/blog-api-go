package handlers

import (
	"blog-api-go/database"
	"blog-api-go/models"
	"blog-api-go/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func GetMe(c *fiber.Ctx) error {
    // ดึง id จาก locals ที่ middleware set ไว้
    userID, ok := c.Locals("userID").(uuid.UUID)
    if !ok {
        return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
    }

    user, err := repository.GetUserByID(userID)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "user not found"})
    }

    return c.JSON(user)
}