package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initDB()
	r.Static("/public/videos", "./public/videos")

	// 基础接口
	getFeed(r)
	user_register(r)
	user_login(r)
	user_getInfo(r)
	post_videos(r)
	user_getList(r)
	r.Run(":9090")
}
