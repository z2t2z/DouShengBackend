package main

import (
	"github.com/gin-gonic/gin"
	"myproj/controller"
)

func main() {
	r := gin.Default()
	controller.InitDB()
	r.Static("/public/videos", "./public/videos")

	// 基础接口
	controller.Get_Feed(r)
	controller.User_Register(r)
	controller.User_Login(r)
	controller.User_GetInfo(r)
	controller.Post_Video(r)
	controller.User_Get_PublishList(r)

	// 互动接口
	controller.User_Like(r)
	controller.Post_Comment(r)
	controller.Get_Favorite_List(r)
	controller.Get_Comment(r)
	
	r.Run(":9090")
}
