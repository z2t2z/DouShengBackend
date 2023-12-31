package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func user_Register(c *gin.Context) {
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
		UserID:     user.ID,
		Token:      token,
	})

}

func user_Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	user, has_user := check_User_Login_Status(token)
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
		UserID:     user.ID,
		Token:      token,
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

func user_GetInfo(c *gin.Context) {
	// id := c.Query("user_id")
	token := c.Query("token")

	user, has_user := check_User_Login_Status(token)

	if !has_user {
		c.JSON(http.StatusOK, UserInfo{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist!",
			User:       DemoUser,
		})
		return
	}

	c.JSON(http.StatusOK, UserInfo{
		StatusCode: 0,
		StatusMsg:  "ok",
		User:       user,
	})
}

func user_Get_PublishList(c *gin.Context) {
	sid := c.PostForm("user_id")
	iid, _ := strconv.ParseInt(sid, 10, 64)

	token := c.Query("token")

	fmt.Println(iid, token)

	user := User{
		Token: token,
	}

	var videoList []Video
	rows, err := db.Table("videos").Select("Play_Url", "Cover_Url", "Favorite_Count", "Comment_Count", "Is_Favorite").Where("user_token = ?", token).Rows()
	if err != nil {
		fmt.Println("select db failed in func: GetList, err:", err)
		return
	}

	for rows.Next() {
		var (
			// ID            int64
			// User_name        string
			Play_Url       string
			Cover_URL      string
			Favorite_Count int64
			Comment_Count  int64
			Is_Favorite    bool
		)

		err := rows.Scan(&Play_Url, &Cover_URL, &Favorite_Count, &Comment_Count, &Is_Favorite)
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
		}

		v := Video{
			// ID:            ID,
			Author:         user,
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

func User_Register(r *gin.Engine) {
	r.POST("/douyin/user/register/", user_Register)
}

func User_Login(r *gin.Engine) {
	r.POST("/douyin/user/login/", user_Login)
}

func User_GetInfo(r *gin.Engine) {
	r.GET("/douyin/user/", user_GetInfo)
}

func User_Get_PublishList(r *gin.Engine) {
	r.GET("/douyin/publish/list/", user_Get_PublishList)
}
