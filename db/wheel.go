package db

import (
	"fmt"
	"go-server/chores/constants"
	"strings"

	"github.com/google/uuid"
)

func InitWheel() {
	myDB.AutoMigrate(&constants.Wheel{})
}

func GetWheelList(openid string) (list []constants.Wheel, err error) {
	dbModel := myDB.Model(&constants.Wheel{}).Where("created_by = ?", openid).Order("created_at DESC")
	err = dbModel.Find(&list).Error
	return list, err
}

func GetWheelDetail(id string) (detail constants.Wheel, err error) {
	err = myDB.Where("id = ?", id).First(&detail).Error
	return detail, err
}

func AddWheel(wheel constants.Wheel) (err error) {
	fmt.Println(wheel)
	wheel.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	err = myDB.Create(&wheel).Error
	return err
}

func EditWheel(wheel constants.Wheel) (err error) {
	id := wheel.ID
	err = myDB.Model(&constants.Wheel{}).Where("id = ?", id).Updates(wheel).Error
	return err
}

func DeleteWheel(id string) (err error) {
	err = myDB.Where("id = ?", id).Delete(&constants.Wheel{}).Error
	return err
}
