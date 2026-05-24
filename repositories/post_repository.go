package repositories

import (
	"blog-api-go/database"
	"blog-api-go/models"
	"database/sql"
	"fmt"
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

func GetPublishedPosts(category, search string, limit, offset int) ([]models.Post, error) {
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

func GetAllPosts(category, search string, limit, offset int) ([]models.Post, error) {
	var posts []models.Post
	err := database.RawDB.Select(&posts, `      
	SELECT 
        posts.*,
        categories.name AS category_name
      FROM posts
      JOIN categories 
        ON posts.category_id = categories.id
      WHERE ($3 = '' OR categories.name = $3)
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

func DeletePostByID(id uint) error {
	_, err := database.RawDB.Exec(`DELETE FROM posts WHERE id = $1`, id)
	return err
}

func CreatePost(post models.CreatePostRequest) error {
	_, err := database.RawDB.Exec(`INSERT INTO posts (title, description, content, category_id, status_id, image) VALUES ($1, $2, $3, $4, $5, $6)`, post.Title, post.Description, post.Content, post.CategoryId, post.StatusId, post.Image)
	return err
}

func UpdatePost(id int, post models.CreatePostRequest) error {
	if id <= 0 {
		return fmt.Errorf("invalid post ID")
	}
	_, err := database.RawDB.Exec(`UPDATE posts SET title = $1, description = $2, content = $3, category_id = $4, status_id = $5, image = $6 WHERE id = $7`, post.Title, post.Description, post.Content, post.CategoryId, post.StatusId, post.Image, id)
	return err
}

func CheckPostExists(postID int) (bool, error) {
    var id int
    err := database.RawDB.Get(&id, "SELECT id FROM posts WHERE id = $1", postID)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, nil  // ไม่เจอ post
        }
        return false, err  // error จริง
    }
    return true, nil  // เจอ post
}
	