package cartservice

import (
	"online-store/models"
	cartrepository "online-store/repositories/cart_repository"
	"online-store/utilities"
)

func CreateCart(request models.Cart) (*models.Cart, error) {
	createCategoryProduct, err := cartrepository.CreateCart(request)
	if err != nil {
		return createCategoryProduct, err
	}
	return createCategoryProduct, nil
}

func GetCart(cartId, productId, userId string) ([]*models.Cart, error) {
	getCart, err := cartrepository.GetCart(cartId, productId, userId)
	if err != nil {
		return getCart, err
	}
	return getCart, nil
}

func DeleteCartById(id string) (int64, error) {
	request := models.Cart{
		DeletedAt: utilities.Times(),
	}
	deleteCart, err := cartrepository.DeleteCartById(id, request.DeletedAt)
	if err != nil {
		return 0, err
	}
	return deleteCart, nil
}

func UpdateCartById(request models.Cart) (int64, error) {
	request.UpdatedAt = utilities.Times()
	deleteCart, err := cartrepository.UpdateCartById(request)
	if err != nil {
		return 0, err
	}
	return deleteCart, nil
}
