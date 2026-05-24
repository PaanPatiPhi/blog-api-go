package handlers

import (
	"blog-api-go/models"
	"blog-api-go/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	postID := c.Params("id")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid post ID",
		})
	}
	exist, err := repositories.CheckPostExists(postIDInt)
	if err != nil {
		return err
	}
	if !exist {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}
	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	comment.PostID = uint(postIDInt)

	err = repositories.CreateComment(comment)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create comment",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Comment created successfully",
	})
}

func GetCommentsByPostID(c *fiber.Ctx) error {
	postID := c.Params("id")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid post ID",
		})
	}
	exist, err := repositories.CheckPostExists(postIDInt)
	if err != nil {
		return err
	}
	if !exist {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}
	comments, err := repositories.GetCommentsByPostID(uint(postIDInt))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get comments",
		})
	}
	return c.JSON(fiber.Map{
    "comments": comments,
})
}
