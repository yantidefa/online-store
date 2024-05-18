package paymentrepository

import (
	"online-store/config"
	"online-store/models"
)

func CreatePayment(request models.Payment) (*models.Payment, error) {
	result := config.DbConn.GormDB.Table("payments").Create(&request)
	if result.Error != nil {
		return nil, result.Error
	}

	return &request, nil
}

func GetPayment(PaymentId, productId, userId string) ([]*models.Payment, error) {
	Payments := []*models.Payment{}
	var queryWhere string
	if productId != "" {
		queryWhere = `p.product_id = '` + productId + `'`
	} else if userId != "" {
		queryWhere = `p.user_id = '` + userId + `'`
	} else if PaymentId != "" {
		queryWhere = ` p.id = '` + PaymentId + `'`
	}
	result := config.DbConn.GormDB.Table("payments p").Joins("JOIN users u ON p.user_id = u.id JOIN products pr ON p.product_id = pr.id").Where("p.deleted_at IS NULL" + queryWhere).Order("p.created_at DESC").Find(&Payments)
	if result.Error != nil {
		return nil, result.Error
	}

	return Payments, nil
}

func DeletePaymentById(id string, deletedAt *string) (int64, error) {
	result := config.DbConn.GormDB.Table("payments").Where("id = ?", id).Update("deleted_at", deletedAt)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}

func UpdatePaymentById(request models.Payment) (int64, error) {
	result := config.DbConn.GormDB.Table("payments").Where("id = ?", request.ID).Updates(request)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}
