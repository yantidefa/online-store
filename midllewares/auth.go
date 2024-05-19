package midllewares

import (
	"log"
	"net/http"
	"online-store/constants"
	usersrepository "online-store/repositories/users_repository"
	"online-store/utilities"

	"github.com/gin-gonic/gin"
)

func AuthenticationCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utilities.SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrEmptyAuthHeader, nil)
			c.Abort()
			return
		}
		err := utilities.ValidateToken(tokenString)
		if err != nil {
			utilities.SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrInvalidToken, err)
			c.Abort()
			return
		}

		getUser, err := usersrepository.GetUser("", "", tokenString)
		if err != nil {
			utilities.SetResponseJSON(c, http.StatusNotFound, nil, constants.ErrLogout, err)
		}

		if getUser.Token == "" {
			utilities.SetResponseJSON(c, http.StatusNotFound, nil, constants.ErrLogout, err)
		}

		jwtPayload, err := utilities.ParseJwtToken(tokenString)
		if err != nil {
			utilities.SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrInvalidToken, err)
			c.Abort()
			return
		}

		if jwtPayload.Role != "Customer" {
			utilities.SetResponseJSON(c, http.StatusUnauthorized, nil, "You Don't Have Permission, Check Your Account", nil)
		}

	}
}
func AuthenticationAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		log.Println(tokenString)
		if tokenString == "" {
			utilities.SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrEmptyAuthHeader, nil)
			c.Abort()
			return
		}
		err := utilities.ValidateToken(tokenString)
		if err != nil {
			utilities.SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrInvalidToken, err)
			c.Abort()
			return
		}

		getUser, err := usersrepository.GetUser("", "", tokenString)
		if err != nil {
			utilities.SetResponseJSON(c, http.StatusNotFound, nil, constants.ErrLogout, err)
		}

		if getUser.Token == "" {
			utilities.SetResponseJSON(c, http.StatusNotFound, nil, constants.ErrLogout, err)
		}

		jwtPayload, err := utilities.ParseJwtToken(tokenString)
		if err != nil {
			utilities.SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrInvalidToken, err)
			c.Abort()
			return
		}

		if jwtPayload.Role != "Admin" {
			utilities.SetResponseJSON(c, http.StatusUnauthorized, nil, "You Don't Have Permission, Check Your Account", nil)
		}

	}
}
