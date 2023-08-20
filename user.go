package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
		return
	} 
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

func getUser(token string) (User, bool) {
	var user User

	result := db.Where("token = ?", token).First(&user)
	if result.RowsAffected == 0 {
		fmt.Println("In func: getUser, found no user")
		return DemoUser, false
	}

	db.Table("users").Where("token = ?", token).Scan(&user)

	return user, true
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	user, has_user := getUser(token)
	if !has_user {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}
	// user := User{
	// 	Name:     username,
	// 	Password: password,
	// 	Token:    token,
	// }
	c.JSON(http.StatusOK, UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "Login successfully",
		Token:      token,
		UserID:     user.ID,
	})

	// result := db.Where("token = ?", token).First(&u)

	// if result.RowsAffected == 0 {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		StatusCode: 1,
	// 		StatusMsg:  "User doesn't exist",
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		StatusCode: 0,
	// 		StatusMsg:  "Login successfully",
	// 		Token:      token,
	// 		UserID:     u.ID,
	// 	})
	// }
}

func GetInfo(c *gin.Context) {
	// id := c.Query("user_id")
	token := c.Query("token")

	user, _ := getUser(token)

	// if err := db.Where("id = ?", id).First(&user); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// db.Where("id = ?", id).First(&user)
	c.JSON(http.StatusOK, UserInfo{
		StatusCode: 0,
		StatusMsg:  "ok",
		User:       user,
	})
}

func GetList(c *gin.Context) {
	sid := c.PostForm("user_id")
	iid, _ := strconv.ParseInt(sid, 10, 64)

	token := c.Query("token")

	fmt.Println(iid, token)

	user := User{
		Token: token,
	}

	// result := db.Where("token = ?", token).First(&user)

	// if result.RowsAffected == 0 {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		StatusCode: 1,
	// 		StatusMsg:  "User doesn't exist",
	// 	})
	// 	return
	// }

	var videoList []Video

	// rows, err := db.Table("videos").Select("Id,  where user_token = ?", token).Rows()
	// rows := db.Query("select id, name, age from users where id= ? ",1)
	// var video Video

	rows, err := db.Table("videos").Select("Author", "Play_Url", "Cover_Url", "Favorite_Count", "Comment_Count", "Is_Favorite").Where("user_token = ?", token).Rows()
	if err != nil {
		fmt.Println("select db failed in func: GetList, err:", err)
		return
	}

	for rows.Next() {
		var (
			// ID            int64
			Author        string
			Play_Url       string
			Cover_URL      string
			Favorite_Count int64
			Comment_Count  int64
			Is_Favorite    bool
		)

		err := rows.Scan(&Author, &Play_Url, &Cover_URL, &Favorite_Count, &Comment_Count, &Is_Favorite)
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
		}

		v := Video{
			// ID:            ID,
			Author:        user,
			Play_URL:       Play_Url,
			Cover_URL:      Cover_URL,
			Favorite_Count: Favorite_Count,
			Comment_Count:  Comment_Count,
			Is_Favorite:    Is_Favorite,
		}
		videoList = append(videoList, v)

	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})

	rows.Close()
}

//执行查询操作

// 	rows,err := Db.Query("SELECT email FROM user_info WHERE user_id>=5")if err !=nil{

// 	fmt.Println("select db failed,err:",err)

// 	return

// 	}// 这里获取的rows是从数据库查的满足user_id>=5的所有行的email信息，rows.Next(),用于循环获取

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
