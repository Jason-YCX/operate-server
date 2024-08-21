package main

import (
	"go-server/db"
	"go-server/router"
	// "go-server/tools"
)

func main() {
	db.InitDB()
	router.InitRouter()
	// tools.InitTools()
}
