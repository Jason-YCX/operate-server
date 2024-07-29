package router

import (
	"encoding/json"
	"fmt"
	"go-server/chores/constants"
	"go-server/chores/utils"
	"go-server/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitWheelRouter() {
	wheelRouter := myRouter.Group("/wheel")

	wheelRouter.GET("/wheel-list", GetWheelList)
	wheelRouter.GET("/wheel-detail", GetWheelDetail)
	wheelRouter.POST("/add-wheel", AddWheel)
	wheelRouter.POST("/edit-wheel", EditWheel)
	wheelRouter.POST("/delete-wheel", DeleteWheel)
}

func GetWheelList(c *gin.Context) {
	openid := utils.GetOpenid(c)
	list, err := db.GetWheelList(openid)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, list)
}

func GetWheelDetail(c *gin.Context) {
	id := c.Query("id")
	wheel, err := db.GetWheelDetail(id)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, wheel)
}

func AddWheel(c *gin.Context) {
	openid := utils.GetOpenid(c)
	var wheel constants.Wheel
	wheel.CreatedBy = openid
	wheel.Options = json.RawMessage(wheel.Options)
	err := c.ShouldBindJSON(&wheel)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.AddWheel(wheel)
	if err != nil {
		fmt.Println("err", err)
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

func EditWheel(c *gin.Context) {
	var wheel constants.Wheel
	err := c.ShouldBindJSON(&wheel)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.EditWheel(wheel)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

func DeleteWheel(c *gin.Context) {
	var opt constants.WheelIdOpt
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.DeleteWheel(opt.Id)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}
