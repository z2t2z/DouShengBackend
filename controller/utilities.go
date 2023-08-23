package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func InitDB() {
	dsn := "root:159357zt@(127.0.0.1:3306)/bytedancedb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	  
	if err != nil {
		panic(err)
	}
}

func check_User_Login_Status(token string) (User, bool) {
	var user User

	result := db.Where("token = ?", token).First(&user)
	if result.RowsAffected == 0 {
		// fmt.Println("In func: getUser, found no user")
		return DemoUser, false
	}

	db.Table("users").Where("token = ?", token).Scan(&user)

	return user, true
}
