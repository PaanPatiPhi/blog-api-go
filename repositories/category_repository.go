package repositories

import (
	"blog-api-go/database"
	"blog-api-go/models"
	"fmt"

	"gorm.io/gorm"
)

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func CreateCategory(category *models.Category) error {
    var existing models.Category
    if err := database.DB.Where("name = ?", category.Name).First(&existing).Error; err == nil {
        return fmt.Errorf("category name already exists")
    }
    return database.DB.Create(&models.Category{Name: category.Name}).Error
}

func UpdateCategory(id uint, category *models.Category) error {
    result := database.DB.Model(&models.Category{}).Where("id = ?", id).Updates(category)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
    return nil
}

func DeleteCategory(id uint) error {
    result := database.DB.Delete(&models.Category{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
    return nil
}
