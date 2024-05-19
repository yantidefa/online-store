package paymenthandler

import (
	"net/http"
	"online-store/constants"
	"online-store/models"
	paymentservice "online-store/services/payment_service"
	"online-store/utilities"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var request models.Payment
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := paymentservice.CreatePayment(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData, err)
	}
}

func GetPaymentByUserId(c *gin.Context) {
	userId := c.Query("user_id")

	data, err := paymentservice.GetPayment(userId)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, err)
	}
}

func DeletePaymentById(c *gin.Context) {
	id := c.Query("id")
	data, err := paymentservice.DeletePaymentById(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, err)
	}
}
