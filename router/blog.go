package router

import (
	"go-server/chores/constants"
	"go-server/chores/utils"
	"go-server/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitBlogRouter() {
	blogRouter := myRouter.Group("/blog")

	blogRouter.GET("/category-list", GetCategoryList)
	blogRouter.POST("/add-category", AddCategory)
	blogRouter.POST("/delete-category", DeleteCategory)

	blogRouter.GET("/tag-list", GetTagList)
	blogRouter.POST("/add-tag", AddTag)
	blogRouter.POST("/delete-tag", DeleteTag)

	blogRouter.GET("/article-list", GetArticleList)
	blogRouter.POST("/add-article", AddArticle)
	blogRouter.POST("/delete-article", DeleteArticle)
}

func GetCategoryList(c *gin.Context) {
	list, err := db.GetCategoryList()
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, list)
}

func AddCategory(c *gin.Context) {
	var category constants.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.AddCategory(category)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

func DeleteCategory(c *gin.Context) {
	var opt constants.IdOpt
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.DeleteCategory(opt.Id)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

func GetTagList(c *gin.Context) {

}

func AddTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}

func GetArticleList(c *gin.Context) {

}

func AddArticle(c *gin.Context) {

}

func DeleteArticle(c *gin.Context) {

}
