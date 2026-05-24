package repositories

import (
	"blog-api-go/database"
	"blog-api-go/models"
)

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
