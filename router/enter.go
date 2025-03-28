package router

import (
	"github.com/gin-gonic/gin"
)

var myGin *gin.Engine
var myRouter *gin.RouterGroup

func InitRouter() {
	myGin = gin.Default()
	myRouter = myGin.Group("/api")

	myGin.Static("/go-static/audio", "./resources/mp3")
	InitBlogRouter()
	InitWheelRouter()
	InitAudioRouter()
	InitLedgerRouter()
	InitWpRouter()

	myGin.Run(":9876")
}
