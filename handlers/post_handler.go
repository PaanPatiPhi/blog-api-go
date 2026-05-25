package handlers

import (
	"blog-api-go/models"
	"blog-api-go/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetPublishedPosts(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 6)
	category := c.Query("category", "")
	search := c.Query("search", "")
	if limit < 1 {
		limit = 1
	}
	if limit > 50 {
		limit = 50
	}
	post, err := repositories.GetPublishedPosts(category, search, limit, (page-1)*limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to get posts: " + err.Error()})
	}
	total, err := repositories.CountPosts(category, search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to get total posts: " + err.Error()})
	}
	totalPages := (total + limit - 1) / limit
	var nextPage *int
	if page < totalPages {
		next := page + 1
		nextPage = &next
	}
	return c.JSON(models.PostResponse{
		Posts:       post,
		TotalPosts:  total,
		TotalPages:  totalPages,
		CurrentPage: page,
		Limit:       limit,
		NextPage:    nextPage,
	})
}

func GetAllPosts(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 6)
	category := c.Query("category", "")
	search := c.Query("search", "")
	if limit < 1 {
		limit = 1
	}
	if limit > 50 {
		limit = 50
	}
	post, err := repositories.GetAllPosts(category, search, limit, (page-1)*limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to get posts: " + err.Error()})
	}
	total, err := repositories.CountAllPosts(category, search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to get total posts: " + err.Error()})
	}
	totalPages := (total + limit - 1) / limit
	var nextPage *int
	if page < totalPages {
		next := page + 1
		nextPage = &next
	}
	return c.JSON(models.PostResponse{
		Posts:       post,
		TotalPosts:  total,
		TotalPages:  totalPages,
		CurrentPage: page,
		Limit:       limit,
		NextPage:    nextPage,
	})
}

func GetPostByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid post ID"})
	}
	post, err := repositories.GetPostByID(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to get post: " + err.Error()})
	}
	return c.JSON(post)
}

func DeletePostByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid post ID"})
	}
	err = repositories.DeletePostByID(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete post: " + err.Error()})
	}
	return c.JSON(fiber.Map{"message": "post deleted successfully"})
}

func CreatePost(c *fiber.Ctx) error {
	var post models.CreatePostRequest
	c.BodyParser(&post)
	validate := validator.New()
	if err := validate.Struct(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	err := repositories.CreatePost(post)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create post: " + err.Error()})
	}
	return c.JSON(fiber.Map{"message": "post created successfully"})
}

func UpdatePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid post ID"})
	}
	var post models.CreatePostRequest
	c.BodyParser(&post)
	validate := validator.New()
	if err := validate.Struct(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	if err := repositories.UpdatePost(id, post);err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update post: " + err.Error()})
	}
	return c.JSON(fiber.Map{"message": "post updated successfully"})
}
