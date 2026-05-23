package repository

import (
	"blog-api-go/database"
	"blog-api-go/models"

	"github.com/google/uuid"
)

func GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, "id = ?", id)
	return &user, result.Error
}
