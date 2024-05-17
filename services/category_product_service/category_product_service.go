package categoryproductservice

import (
	"online-store/models"
	categoryproductrepository "online-store/repositories/category_product_repository"
	"online-store/utilities"
)

func CreateCategoryProduct(request models.CategoryProduct) (*models.CategoryProduct, error) {
	createCategoryProduct, err := categoryproductrepository.CreateCategoryProduct(request)
	if err != nil {
		return createCategoryProduct, err
	}
	return createCategoryProduct, nil
}

func GetAllCategoryProduct() (*[]models.CategoryProduct, error) {
	getCategoryProduct, err := categoryproductrepository.GetAllCategoryProduct()
	if err != nil {
		return getCategoryProduct, err
	}
	return getCategoryProduct, nil
}

func GetCategoryProductById(id string) (*models.CategoryProduct, error) {
	getCategoryProduct, err := categoryproductrepository.GetCategoryProductById(id)
	if err != nil {
		return getCategoryProduct, err
	}

	return getCategoryProduct, nil
}

func DeleteCategoryProductById(id string) (int64, error) {
	request := models.CategoryProduct{
		DeletedAt: utilities.Times(),
	}
	deleteCategoryProduct, err := categoryproductrepository.DeleteCategoryProductById(id, request.DeletedAt)
	if err != nil {
		return 0, err
	}
	return deleteCategoryProduct, nil
}

func UpdateCategoryProductById(request models.CategoryProduct) (int64, error) {
	request.UpdatedAt = utilities.Times()
	deleteCategoryProduct, err := categoryproductrepository.UpdateCategoryProductById(request)
	if err != nil {
		return 0, err
	}
	return deleteCategoryProduct, nil
}
