package repositories

import (
	"blog-api-go/database"
)

func GetLikesCount(postID uint) (int, error) {
	var count int
	err := database.RawDB.Get(&count, "SELECT COUNT(*) FROM likes WHERE post_id = $1", postID)
	return count, err
}

func GetUserLikeStatus(postID uint, userID string) (bool, error) {
	var liked bool
	err := database.RawDB.Get(&liked, "SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = $1 AND user_id = $2) as liked", postID, userID)
	return liked, err
}

func DeleteLike(postID uint, userID string) error {
	_, err := database.RawDB.Exec("DELETE FROM likes WHERE post_id = $1 AND user_id = $2", postID, userID)
	return err
}

func CreateLike(postID uint, userID string) error {
	_, err := database.RawDB.Exec("INSERT INTO likes (post_id, user_id) VALUES ($1, $2)", postID, userID)
	return err
}
