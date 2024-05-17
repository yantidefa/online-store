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
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := productservice.CreateProduct(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData, err)
	}
}

func GetAllProduct(c *gin.Context) {
	categoryId := c.Query("category_id")
	categoryName := c.Query("category_name")

	data, err := productservice.GetAllProduct(categoryId, categoryName)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, err)
	}
}

func GetProductById(c *gin.Context) {
	id := c.Query("id")
	data, err := productservice.GetProductById(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedDataDeleted, err)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, err)
	}
}

func DeleteProductById(c *gin.Context) {
	id := c.Query("id")
	data, err := productservice.DeleteProductById(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, err)
	}
}

func UpdateProductById(c *gin.Context) {
	var request models.Product
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := productservice.UpdateProductById(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData, err)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessUpdateData, err)
	}
}
