package productrepository

import (
	"online-store/config"
	"online-store/models"
	"online-store/utilities"
)

func CreateProduct(request models.Product) (*models.Product, error) {
	request.CreatedAt = utilities.Times()
	result := config.DbConn.GormDB.Table("products").Create(&request)
	if result.Error != nil {
		return nil, result.Error
	}

	return &request, nil
}

func GetProduct(productId, categoryId, categoryName string) ([]*models.Product, error) {
	products := []*models.Product{}
	var queryWhere string
	if categoryId != "" {
		queryWhere = `category_id = '` + categoryId + `'`
	} else if categoryName != "" {
		queryWhere = `c.name = '` + categoryName + `'`
	} else if productId != "" {
		queryWhere =` p.id = '` + productId + `'`
	}
	result := config.DbConn.GormDB.Table("products p").Joins("JOIN categories c ON p.category_id = c.id").Where(queryWhere).Where("p.deleted_at IS NULL").Order("p.created_at DESC").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func DeleteProductById(id string, deletedAt *string) (int64, error) {
	result := config.DbConn.GormDB.Table("products").Where("id = ?", id).Update("deleted_at", deletedAt)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}

func UpdateProductById(request models.Product) (int64, error) {
	result := config.DbConn.GormDB.Table("products").Where("id = ?", request.ID).Updates(request)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}
