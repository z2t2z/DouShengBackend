package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func CheckUserLoginStatus(c *gin.Context) bool {
	token := c.PostForm("token")

	user := User{
		Token: token,
	}

	result := db.Where("token = ?", token).First(&user)

	return result.RowsAffected != 0
}

func publishAction(c *gin.Context) {
	token := c.PostForm("token")

	user, has_user := getUser(token)

	// result := db.Where("token = ?", token).First(&user)

	if !has_user {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}

	// user := User{
	// 	Token: token,
	// }

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%s_%s", token, filename)
	saveFile := filepath.Join("./public/videos", finalName)
	// fmt.Println(saveFile)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// user := User{
	// 	Token: token,
	// }
	// var user User
	user.Token = token

	title := c.Query("title")
	Play_URL := "http://192.168.0.105:9090/public/videos" + finalName
	video := Video{
		// User_token: token,
		Play_URL:    Play_URL,
		Title:      title,
	}

	// db.Select("PlayURL", "User_token", "Title").Create(&video)
	if err := db.Select("Play_URL", "User_token", "Title").Create(&video).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// func PublishList(c *gin.Context) {
// 	token := c.PostForm("token")
// 	fmt.Println(token)
// 	c.JSON(http.StatusOK, VideoListResponse{
// 		Response: Response{
// 			StatusCode: 0,
// 		},
// 		VideoList: DemoVideos,
// 	})
// }

func post_videos(r *gin.Engine) {
	r.POST("/douyin/publish/action/", publishAction)
}

// func get_PublishList(r *gin.Engine) {
// 	r.GET("/douyin/publish/list/", PublishList)
// }
