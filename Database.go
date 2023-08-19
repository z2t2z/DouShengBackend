package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "root:159357zt@(127.0.0.1:3306)/bytedancedb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}
}
