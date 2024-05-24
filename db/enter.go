package db

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var myDB *gorm.DB

func InitDB() {
	// 设置东8区时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = loc
	// 数据库连接信息
	dsn := "root:Jason123....@tcp(127.0.0.1:3306)/JASON_SERVER?parseTime=true&loc=Local"

	// 打开数据库连接
	myDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	myDB.Logger = myDB.Logger.LogMode(4)

	InitBlog()
}
