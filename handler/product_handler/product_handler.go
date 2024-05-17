package producthandler

import (
	"net/http"
	"online-store/constants"
	"online-store/models"
	productservice "online-store/services/product_service"
	"online-store/utilities"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var request models.Product
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData)
		return
	}

	data, err := productservice.CreateProduct(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData)
	}
}

func GetAllProduct(c *gin.Context) {
	categoryId := c.Query("category_id")
	categoryName := c.Query("category_name")

	data, err := productservice.GetAllProduct(categoryId, categoryName)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData)
	}
}

func GetProductById(c *gin.Context) {
	id := c.Query("id")
	data, err := productservice.GetProductById(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedDataDeleted)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData)
	}
}

func DeleteProductById(c *gin.Context) {
	id := c.Query("id")
	data, err := productservice.DeleteProductById(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData)
	}
}

func UpdateProductById(c *gin.Context) {
	var request models.Product
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData)
		return
	}

	data, err := productservice.UpdateProductById(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessUpdateData)
	}
}
