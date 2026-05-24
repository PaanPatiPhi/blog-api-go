package repositories

import (
	"blog-api-go/database"
	"blog-api-go/models"
)
func CountPosts(category, search string) (int, error) {
	var count int
	err := database.RawDB.Get(&count, `      
	SELECT COUNT(*) 
	FROM posts
	JOIN categories 
		ON posts.category_id = categories.id
	WHERE posts.status_id = 2
	AND ($1 = '' OR categories.name = $1)
	AND ($2 = '' OR posts.title ILIKE '%' || $2 || '%')`, category, search)
	return count, err
}

func GetPosts(category, search string, limit, offset int) ([]models.Post, error) {
	var posts []models.Post
	err := database.RawDB.Select(&posts, `      
	SELECT 
        posts.*,
        categories.name AS category_name
      FROM posts
      JOIN categories 
        ON posts.category_id = categories.id
      WHERE posts.status_id = 2
      AND ($3 = '' OR categories.name = $3)
      AND ($4 = '' OR posts.title ILIKE '%' || $4 || '%')	
      ORDER BY posts.id DESC
      LIMIT $1 OFFSET $2`, limit, offset, category, search)
	return posts, err
}

func GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := database.RawDB.Get(&post, `      
	SELECT 
        posts.*, categories.name AS category_name
      FROM posts
	  JOIN categories 
		ON posts.category_id = categories.id
      WHERE posts.id = $1`, id)
	return &post, err
}
