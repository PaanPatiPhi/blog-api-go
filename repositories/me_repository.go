package repositories

import (
	"blog-api-go/database"
	"blog-api-go/models"
)

func GetMe(userID string) (*models.Me, error) {
	var me models.Me
	err := database.RawDB.Get(&me, `SELECT id, username, name, profile_pic, role FROM users WHERE id = $1`, userID)
	return &me, err
}

func UpdateUserProfile(userID string, username *string, name *string, profilePic *string) error {
    _, err := database.RawDB.Exec(`
        UPDATE users
        SET 
            username    = COALESCE($1, username),
            name        = COALESCE($2, name),
            profile_pic = COALESCE($3, profile_pic)
        WHERE id = $4
    `, username, name, profilePic, userID)
    return err
}