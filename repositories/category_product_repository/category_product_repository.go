package categoryproductrepository

import (
	"online-store/config"
	"online-store/models"
	"online-store/utilities"
)

func CreateCategoryProduct(request models.CategoryProduct) (*models.CategoryProduct, error) {
	request.CreatedAt = utilities.Times()
	result := config.DbConn.GormDB.Table("categories").Create(&request)
	if result.Error != nil {
		return nil, result.Error
	}

	return &request, nil
}

func GetAllCategoryProduct() (*[]models.CategoryProduct, error) {
	categories := []models.CategoryProduct{}
	result := config.DbConn.GormDB.Table("categories").Where("deleted_at IS NULL").Order("created_at DESC").Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return &categories, nil
}

func GetCategoryProductById(id string) (*models.CategoryProduct, error) {
	var category models.CategoryProduct
	result := config.DbConn.GormDB.Table("categories").Where("id = ? AND deleted_at IS NULL", id).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func DeleteCategoryProductById(id string, deletedAt *string) (int64, error) {
	result := config.DbConn.GormDB.Table("categories").Where("id = ?", id).Update("deleted_at", deletedAt)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}

func UpdateCategoryProductById(request models.CategoryProduct) (int64, error) {
	result := config.DbConn.GormDB.Table("categories").Where("id = ?", request.ID).Updates(request)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}
