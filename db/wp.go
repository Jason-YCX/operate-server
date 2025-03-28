package db

import (
	"go-server/chores/constants"
	"strings"

	"github.com/google/uuid"
)

func InitWp() {
	myDB.AutoMigrate(&constants.WpResourceType{})
	myDB.AutoMigrate(&constants.WpResourceDetail{})
}

func GetWpResourceTypeList() (list []constants.WpResourceType, err error) {
	dbModel := myDB.Model(&constants.WpResourceType{}).Order("created_at DESC")
	err = dbModel.Find(&list).Error
	return list, err
}

func AddWpResourceType(wpResourceType constants.WpResourceType) (err error) {
	wpResourceType.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	err = myDB.Create(&wpResourceType).Error
	return err
}

func UpdateWpResourceType(wpResourceType constants.WpResourceType) (err error) {
	err = myDB.Model(&constants.WpResourceType{}).Where("id=?", wpResourceType.ID).Updates(wpResourceType).Error
	return err
}

func DeleteWpResourceType(id string) (err error) {
	err = myDB.Where("id=?", id).Delete(&constants.WpResourceType{}).Error
	return err
}

func GetWpResourceList() (list []constants.WpResourceDetail, err error) {
	dbModel := myDB.Model(&constants.WpResourceDetail{}).Order("created_at DESC")
	err = dbModel.Find(&list).Error
	return list, err
}

func AddWpResource(WpResource constants.WpResourceDetail) (err error) {
	WpResource.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	err = myDB.Create(&WpResource).Error
	return err
}

func UpdateWpResource(WpResource constants.WpResourceDetail) (err error) {
	err = myDB.Model(&constants.WpResourceDetail{}).Where("id=?", WpResource.ID).Updates(WpResource).Error
	return err
}

func DeleteWpResource(id string) (err error) {
	err = myDB.Where("id=?", id).Delete(&constants.WpResourceDetail{}).Error
	return err
}
