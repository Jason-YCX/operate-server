package router

import (
	"fmt"
	"go-server/chores/constants"
	"go-server/chores/utils"
	"go-server/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitLedgerRouter() {
	ledgerRouter := myRouter.Group("/ledger")

	ledgerRouter.POST("ledger-list", GetLedgerList)
	ledgerRouter.POST("add-ledger", AddLedger)
	ledgerRouter.PUT("edit-ledger", EditLedger)
}

func GetLedgerList(c *gin.Context) {
	var opt constants.LedgerListOpt
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		fmt.Println(err.Error())
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	list, err := db.GetLedgerList(opt)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, list)
}

func AddLedger(c *gin.Context) {
	var ledger constants.Ledger
	err := c.ShouldBindJSON(&ledger)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.AddLedger(ledger)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}

func EditLedger(c *gin.Context) {
	var ledger constants.Ledger
	err := c.ShouldBindJSON(&ledger)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	err = db.EditLedger(ledger)
	if err != nil {
		utils.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}
	utils.SendResponse(c, http.StatusOK, 200, nil)
}
