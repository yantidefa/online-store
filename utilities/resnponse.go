package utilities

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

func SetResponseJSON(c *gin.Context, code int, data interface{}, msg string) {

	response := Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}

	c.JSON(code,
		response)
	c.Abort()
}
