package handlers

import (
	"blog-api-go/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetLikesCount(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid post ID",
		})
	}
	count, err := repositories.GetLikesCount(uint(postId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get likes count",
		})
	}
	return c.JSON(fiber.Map{
		"count": count,
	})
}

func GetUserLikeStatus(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid post ID",
		})
	}
	userId := c.Locals("userID").(string)
	liked, err := repositories.GetUserLikeStatus(uint(postId), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user like status",
		})
	}
	return c.JSON(fiber.Map{
		"liked": liked,
	})
}

func ToggleLike(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid post ID",
		})
	}
	exist, err := repositories.CheckPostExists(postId)
	if err != nil {
		return err
	}
	if !exist {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}
	userId := c.Locals("userID").(string)
	liked, err := repositories.GetUserLikeStatus(uint(postId), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user like status",
		})
	}
	if liked {
		err = repositories.DeleteLike(uint(postId), userId)
	} else {
		err = repositories.CreateLike(uint(postId), userId)
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to toggle like",
		})
	}
	return c.JSON(fiber.Map{
		"liked": !liked,
	})
}