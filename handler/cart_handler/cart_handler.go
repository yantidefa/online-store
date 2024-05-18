package carthandler

import (
	"net/http"
	"online-store/constants"
	"online-store/models"
	cartservice "online-store/services/cart_service"
	"online-store/utilities"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	var request models.Cart
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := cartservice.CreateCart(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData, err)
	}
}

func GetCart(c *gin.Context) {
	userId := c.Query("user_id")
	ProductId := c.Query("product_id")
	CartId := c.Query("Cart_id")

	data, err := cartservice.GetCart(CartId, ProductId, userId)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, err)
	}
}

func DeleteCartById(c *gin.Context) {
	id := c.Query("id")
	data, err := cartservice.DeleteCartById(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, err)
	}
}

func UpdateCartById(c *gin.Context) {
	var request models.Cart
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := cartservice.UpdateCartById(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData, err)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessUpdateData, err)
	}
}
