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

func Action(c *gin.Context) {
	token := c.PostForm("token")

	user := User{
		Token: token,
	}

	result := db.Where("token = ?", token).First(&user)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.ID, filename)
	saveFile := filepath.Join("./public/", finalName)

	fmt.Println(saveFile)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func PublishList(c *gin.Context) {
	token := c.PostForm("token")
	fmt.Println(token)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}

func post_videos(r *gin.Engine) {
	r.POST("/douyin/publish/action/", Action)
}

// func get_PublishList(r *gin.Engine) {
// 	r.GET("/douyin/publish/list/", PublishList)
// }
