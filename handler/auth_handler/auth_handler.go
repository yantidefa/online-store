package authhandler

import (
	"net/http"
	"online-store/constants"
	"online-store/models"
	authservice "online-store/services/auth_service"
	"online-store/utilities"

	"github.com/gin-gonic/gin"
)

func RegisterCustomer(c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := authservice.RegisterCustomer(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData, err)
	}
}
func RegisterAdmin(c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := authservice.RegisterAdmin(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData, err)
	}
}
func Login(c *gin.Context) {
	var request models.Login
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := authservice.Login(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	} else {
		utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData, err)
	}
}
