package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetHeaders(c *gin.Context) {
	// 获取所有请求头信息
	headers := c.Request.Header
	for key, value := range headers {
		fmt.Println(key, value)
	}
}

func GetOpenid(c *gin.Context) (openid string) {
	// 获取单个请求头信息
	Authorization := c.GetHeader("Authorization")
	openid = strings.TrimPrefix(Authorization, "Bearer ")
	return openid
}
