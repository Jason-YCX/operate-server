package main

import (
	"go-server/db"
	"go-server/router"
)

func main() {
	db.InitDB()
	router.InitRouter()
}
