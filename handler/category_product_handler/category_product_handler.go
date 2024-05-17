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
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData)
		return
	}

	data, err := categoryproductservice.CreateCategoryProduct(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData)
	}
}

func GetCategoryProductAll(c *gin.Context) {
	data, err := categoryproductservice.GetAllCategoryProduct()
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData)
	}
}

func GetCategoryProductById(c *gin.Context) {
	id := c.Query("id")
	data, err := categoryproductservice.GetCategoryProductById(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedDataDeleted)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData)
	}
}

func DeleteCategoryProductById(c *gin.Context) {
	id := c.Query("id")
	data, err := categoryproductservice.DeleteCategoryProductById(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData)
	}
}

func UpdateCategoryProductById(c *gin.Context) {
	var request models.CategoryProduct
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData)
		return
	}

	data, err := categoryproductservice.UpdateCategoryProductById(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessUpdateData)
	}
}
