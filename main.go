package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 基于Gin框架的极简版抖音开发
	r := gin.Default()
	initDB()
	r.Static("/public", "./public")

	// 基础接口
	getFeed(r)
	user_register(r)
	user_login(r)
	user_getInfo(r)
	post_videos(r)
	user_getList(r)
	// get_PublishList(r)
	// abcde
	r.Run(":9090")
}
