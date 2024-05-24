package utils

import (
	"github.com/gin-gonic/gin"
)

var codeMessage = map[int]string{
	200: "Success",
	400: "Bad Request",
	404: "Not Found",
	500: "Internal Server Error",
}

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, httpStatus int, code int, data interface{}) {
	message := codeMessage[code]
	if message == "" {
		message = "Internal Server Error"
	}
	c.JSON(httpStatus, ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
