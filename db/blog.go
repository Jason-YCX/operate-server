package db

import (
	"go-server/chores/constants"
	"strings"

	"github.com/google/uuid"
)

func InitBlog() {
	myDB.AutoMigrate(&constants.Category{})
	myDB.AutoMigrate(&constants.Tag{})
	myDB.AutoMigrate(&constants.Note{})
}

// 分类
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

func EditCategory(category constants.Category) (err error) {
	id := category.ID
	err = myDB.Model(&constants.Category{}).Where("id = ?", id).Updates(category).Error
	return err
}

func DeleteCategory(id string) (err error) {
	err = myDB.Where("id = ?", id).Delete(&constants.Category{}).Error
	return err
}

// 标签
func GetTagList() (list []constants.Tag, err error) {
	dbModel := myDB.Model(&constants.Tag{}).Order("created_at DESC")
	err = dbModel.Find(&list).Error
	return list, err
}

func AddTag(tag constants.Tag) (err error) {
	tag.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	err = myDB.Create(&tag).Error
	return err
}

func EditTag(tag constants.Tag) (err error) {
	id := tag.ID
	err = myDB.Model(&constants.Tag{}).Where("id = ?", id).Updates(tag).Error
	return err
}

func DeleteTag(id string) (err error) {
	err = myDB.Where("id = ?", id).Delete(&constants.Tag{}).Error
	return err
}

// 笔记
func GetNoteList(opt constants.NoteListOpt) (list []constants.Note, err error) {
	dbModel := myDB.Model(&constants.Note{}).Order("created_at DESC")
	if opt.Keyword != "" {
		keyword := "%" + opt.Keyword + "%"
		dbModel = dbModel.Where("title LIKE ? OR content LIKE ?", keyword, keyword)
	}
	err = dbModel.Find(&list).Error
	return list, err
}

func AddNote(note constants.Note) (err error) {
	note.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	err = myDB.Create(&note).Error
	return err
}

func EditNote(note constants.Note) (err error) {
	id := note.ID
	err = myDB.Model(&constants.Note{}).Where("id = ?", id).Updates(note).Error
	return err
}

func DeleteNote(id string) (err error) {
	err = myDB.Where("id = ?", id).Delete(&constants.Note{}).Error
	return err
}
