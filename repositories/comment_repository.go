package repositories

import (
	"blog-api-go/database"
	"blog-api-go/models"
)

func CreateComment(comment models.Comment) error {
	_, err := database.RawDB.Exec(`INSERT INTO comments (post_id, user_id, comment_text) VALUES ($1, $2, $3)`, comment.PostID, comment.UserID, comment.CommentText)
	return err
}

func GetCommentsByPostID(postID uint) ([]models.CommentWithUser, error) {
    var comments []models.CommentWithUser
    err := database.RawDB.Select(&comments, `
      SELECT c.*, u.username, u.profile_pic
      FROM comments c
      LEFT JOIN users u ON c.user_id = u.id
      WHERE c.post_id = $1
      ORDER BY c.created_at DESC
    `, postID)
    return comments, err
}