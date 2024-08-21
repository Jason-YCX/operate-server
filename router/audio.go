package router

import (
	"go-server/chores/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitAudioRouter() {
	audioRouter := myRouter.Group("/audio")

	audioRouter.GET("/audio-list", getAudioList)
}

func getAudioList(c *gin.Context) {
	audioDir := "./resources/mp3/"
	files, err := os.ReadDir(audioDir)
	if err != nil {
		utils.SendResponse(c, http.StatusOK, 400, nil)
		return
	}
	audioList := []string{}
	for _, file := range files {

		audioList = append(audioList, file.Name())
	}
	utils.SendResponse(c, http.StatusOK, 200, audioList)
}
