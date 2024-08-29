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
	blogRouter.POST("/edit-category", EditCategory)
	blogRouter.POST("/delete-category", DeleteCategory)

	blogRouter.GET("/tag-list", GetTagList)
	blogRouter.POST("/add-tag", AddTag)
	blogRouter.POST("/edit-tag", EditTag)
	blogRouter.POST("/delete-tag", DeleteTag)

	blogRouter.POST("/note-list", GetNoteList)
	blogRouter.POST("/add-note", AddNote)
	blogRouter.POST("/edit-note", EditNote)
	blogRouter.POST("/delete-note", DeleteNote)
}

// 分类
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

func EditCategory(c *gin.Context) {
	var category constants.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.EditCategory(category)
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

// 标签
func GetTagList(c *gin.Context) {
	list, err := db.GetTagList()
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, list)
}

func AddTag(c *gin.Context) {
	var tag constants.Tag
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.AddTag(tag)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

func EditTag(c *gin.Context) {
	var tag constants.Tag
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.EditTag(tag)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

func DeleteTag(c *gin.Context) {
	var opt constants.IdOpt
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.DeleteTag(opt.Id)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

// 笔记
func GetNoteList(c *gin.Context) {
	var opt constants.NoteListOpt
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	list, err := db.GetNoteList(opt)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, list)
}

func AddNote(c *gin.Context) {
	var note constants.Note
	err := c.ShouldBindJSON(&note)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.AddNote(note)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

func EditNote(c *gin.Context) {
	var note constants.Note
	err := c.ShouldBindJSON(&note)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.EditNote(note)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

func DeleteNote(c *gin.Context) {
	var opt constants.IdOpt
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.DeleteNote(opt.Id)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}
