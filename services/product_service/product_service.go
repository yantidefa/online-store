package productservice

import (
	"online-store/models"
	productrepository "online-store/repositories/product_repository"
	"online-store/utilities"
)

func CreateProduct(request models.Product) (*models.Product, error) {
	createProduct, err := productrepository.CreateProduct(request)
	if err != nil {
		return createProduct, err
	}
	return createProduct, nil
}

func GetProduct(productId, categoryId, categoryName string) ([]*models.Product, error) {
	getProduct, err := productrepository.GetProduct(productId, categoryId, categoryName)
	if err != nil {
		return getProduct, err
	}
	return getProduct, nil
}

func DeleteProductById(id string) (int64, error) {
	request := models.Product{
		DeletedAt: utilities.Times(),
	}
	deleteProduct, err := productrepository.DeleteProductById(id, request.DeletedAt)
	if err != nil {
		return 0, err
	}
	return deleteProduct, nil
}

func UpdateProductById(request models.Product) (int64, error) {
	request.UpdatedAt = utilities.Times()
	deleteProduct, err := productrepository.UpdateProductById(request)
	if err != nil {
		return 0, err
	}
	return deleteProduct, nil
}
