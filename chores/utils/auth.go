package utils

import "github.com/gin-gonic/gin"

func AuthRoute() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
