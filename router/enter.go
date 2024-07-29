package router

import (
	"github.com/gin-gonic/gin"
)

var myGin *gin.Engine
var myRouter *gin.RouterGroup

func InitRouter() {
	myGin = gin.Default()
	myRouter = myGin.Group("/api")
	InitBlogRouter()
	InitWheelRouter()

	myGin.Run(":9876")
}
