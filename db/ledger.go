package db

import (
	"go-server/chores/constants"
)

func InitLedger() {
	myDB.AutoMigrate(&constants.Ledger{})
}

func GetLedgerList(opt constants.LedgerListOpt) (list []constants.Ledger, err error) {
	dbModel := myDB.Model(&constants.Ledger{}).Order("created_at desc")
	err = dbModel.Find(&list).Error
	return list, err
}

func AddLedger(ledger constants.Ledger) (err error) {
	err = myDB.Create(&ledger).Error
	return err
}

func EditLedger(ledger constants.Ledger) (err error) {
	id := ledger.ID
	err = myDB.Model(&constants.Ledger{}).Where("id=?", id).First(ledger).Updates(&ledger).Error
	return err
}
