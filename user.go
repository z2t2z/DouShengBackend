package main

import (
	"fmt"
	"net/http"
	"time"
	"strconv"  
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	var user User
	// First, Find, Where
	result := db.Find(&user, "token = ?", token)

	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "User already exists",
		})
	} else {
		// 先插入数据库
		u := User{
			Name:       username,
			Password:   password,
			Token:      token,
			CreateDate: time.Now(),
		}

		// 要么传指针, 要么传赋了非零值的id
		if err := db.Create(&u).Error; err != nil {
			fmt.Println("Error In function: Register, Insert failed!\n", err)
			return
		}

		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 0,
			StatusMsg:  "Created a user successfully",
			Token:      token,
			UserID:     1,
		})

	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	u := User{
		Name:     username,
		Password: password,
		Token:    token,
	}

	result := db.Where("token = ?", token).First(&u)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 0,
			StatusMsg:  "Login successfully",
			Token:      token,
			UserID:     u.ID,
		})
	}
}

func GetInfo(c *gin.Context) {
	id := c.Query("user_id")
	// token := c.Query("token")

	u := User{}

	// 好像return了
	// if err := db.Where("id = ?", id).First(&u); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	db.Where("id = ?", id).First(&u)
	c.JSON(http.StatusOK, UserInfo{
		StatusCode: 0,
		StatusMsg:  "ok",
		User:       u,
	})
}

func GetList(c *gin.Context) {
	sid := c.Query("user_id")
	iid, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		fmt.Println("Stoi error", err)

	}
	token := c.Query("token")
	user := User{
		ID:    iid,
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

	// for temporary use 
	type VideoListResponse struct {
		Response
		VideoList []Video `json:"video_list"`
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}

func user_register(r *gin.Engine) {
	r.POST("/douyin/user/register/", Register)
}

func user_login(r *gin.Engine) {
	r.POST("/douyin/user/login/", Login)
}

func user_getInfo(r *gin.Engine) {
	r.GET("/douyin/user/", GetInfo)
}

func user_getList(r *gin.Engine) {
	r.GET("/douyin/publish/list/", GetList)
}
