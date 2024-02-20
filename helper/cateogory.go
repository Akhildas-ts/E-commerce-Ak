package helper

import (
	"ak/database"
	"ak/models"
)

func CheckCategoryId(CategoryID int) (error) {
	var count int
	errRes := database.DB.Raw("SELECT COUNT(*) FROM categories WHERE id = ?", CategoryID).Scan(&count).Error

	if errRes != nil {
		return errRes
	}

	if count == 0 {
		return models.ThereIsNOCategory
	}

	return nil
}
