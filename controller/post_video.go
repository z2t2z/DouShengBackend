package controller

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

func post_Video(c *gin.Context) {
	token := c.PostForm("token")

	user, has_user := check_User_Login_Status(token)

	if !has_user {
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

	user.Token = token

	title := c.PostForm("title")
	Play_URL := "http://192.168.0.105:9090/public/videos/" + finalName
	video := Video{
		User_id:    user.ID,
		Author:     user,
		Play_URL:   Play_URL,
		Cover_URL:  "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		Title:      title,
	}

	// db.Select("PlayURL", "User_token", "Title").Create(&video)
	if err := db.Select("User_id", "Play_URL", "Cover_URL", "Title").Create(&video).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func Post_Video(r *gin.Engine) {
	r.POST("/douyin/publish/action/", post_Video)
}
