package cartrepository

import (
	"online-store/config"
	"online-store/models"
	"online-store/utilities"
)

func CreateCart(request models.Cart) (*models.Cart, error) {
	request.CreatedAt = utilities.Times()
	result := config.DbConn.GormDB.Table("carts").Create(&request)
	if result.Error != nil {
		return nil, result.Error
	}

	return &request, nil
}

func GetCart(cartId, productId, userId string) ([]*models.Cart, error) {
	carts := []*models.Cart{}
	var queryWhere string
	if cartId != "" {
		queryWhere = `AND id = '` + cartId + `'`
	} else if userId != "" {
		queryWhere = `AND user_id = '` + userId + `'`
	} else if productId != "" {
		queryWhere = `AND product_id = '` + productId + `'`
	}
	result := config.DbConn.GormDB.Table("carts").Where("deleted_at IS NULL " + queryWhere).Order("created_at DESC").Find(&carts)
	if result.Error != nil {
		return nil, result.Error
	}

	return carts, nil
}

func DeleteCartById(id string, deletedAt *string) (int64, error) {
	result := config.DbConn.GormDB.Table("carts").Where("id = ?", id).Update("deleted_at", deletedAt)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}

func UpdateCartById(request models.Cart) (int64, error) {
	result := config.DbConn.GormDB.Table("carts").Where("id = ?", request.ID).Updates(request)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}
