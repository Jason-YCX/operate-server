package router

import (
	"go-server/chores/constants"
	"go-server/chores/utils"
	"go-server/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitWpRouter() {
	wpRouter := myRouter.Group("/wp")

	wpRouter.GET("/resource-type-list", GetResourceTypeList)
	wpRouter.POST("/add-resource-type", AddResourceType)
	wpRouter.GET("/resource-list", GetResourceList)
	wpRouter.POST("/add-resource", AddResource)
	wpRouter.POST("/delete-resource", DeleteResource)
}

// 资源类型列表
func GetResourceTypeList(c *gin.Context) {
	list, err := db.GetWpResourceTypeList()
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, list)
}

// 新增资源类型
func AddResourceType(c *gin.Context) {
	var resourceType constants.WpResourceType
	err := c.BindJSON(&resourceType)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.AddWpResourceType(resourceType)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, resourceType)
}

// 删除资源类型
func DeleteResourceType(c *gin.Context) {
	var opt constants.WpResourceTypeIdOpt
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.DeleteWpResourceType(opt.ID)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

// 更新资源类型
func UpdateResourceType(c *gin.Context) {
	var resourceType constants.WpResourceType
	err := c.BindJSON(&resourceType)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.UpdateWpResourceType(resourceType)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

// 获取资源列表
func GetResourceList(c *gin.Context) {
	list, err := db.GetWpResourceList()
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, list)
}

// 新增资源
func AddResource(c *gin.Context) {
	var resource constants.WpResourceDetail
	err := c.BindJSON(&resource)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.AddWpResource(resource)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, resource)
}

// 删除资源
func DeleteResource(c *gin.Context) {
	var opt constants.WpResourceIdOpt
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.DeleteWpResource(opt.ID)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

// 更新资源
func UpdateResource(c *gin.Context) {
	var resource constants.WpResourceDetail
	err := c.BindJSON(&resource)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.UpdateWpResource(resource)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}
