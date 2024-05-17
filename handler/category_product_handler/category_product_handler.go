package categoryproducthandler

import (
	"net/http"
	"online-store/constants"
	"online-store/models"
	categoryproductservice "online-store/services/category_product_service"
	"online-store/utilities"

	"github.com/gin-gonic/gin"
)

func CreateCategoryProduct(c *gin.Context) {
	var request models.CategoryProduct
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := categoryproductservice.CreateCategoryProduct(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData, err)
	}
}

func GetCategoryProductAll(c *gin.Context) {
	data, err := categoryproductservice.GetAllCategoryProduct()
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, err)
	}
}

func GetCategoryProductById(c *gin.Context) {
	id := c.Query("id")
	data, err := categoryproductservice.GetCategoryProductById(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedDataDeleted, err)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, err)
	}
}

func DeleteCategoryProductById(c *gin.Context) {
	id := c.Query("id")
	data, err := categoryproductservice.DeleteCategoryProductById(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, err)
	}
}

func UpdateCategoryProductById(c *gin.Context) {
	var request models.CategoryProduct
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := categoryproductservice.UpdateCategoryProductById(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData, err)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessUpdateData, err)
	}
}
