package db

import (
	"go-server/chores/constants"
	"strings"

	"github.com/google/uuid"
)

func InitBlog() {
	myDB.AutoMigrate(&constants.Category{})
	myDB.AutoMigrate(&constants.Tag{})
	myDB.AutoMigrate(&constants.Article{})
}

func GetCategoryList() (list []constants.Category, err error) {
	dbModel := myDB.Model(&constants.Category{}).Order("created_at DESC")
	err = dbModel.Find(&list).Error
	return list, err
}

func AddCategory(category constants.Category) (err error) {
	category.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	err = myDB.Create(&category).Error
	return err
}

func DeleteCategory(id string) (err error) {
	err = myDB.Where("id = ?", id).Delete(&constants.Category{}).Error
	return err
}

func EditCategory(category constants.Category) (err error) {
	return
}
