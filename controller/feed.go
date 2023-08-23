package controller

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	// util "myproj/controller/utility"
)

// 这里没有限制返回的具体时间戳
func get_Feed(c *gin.Context) {
	var videoList []Video

	rows, err := db.Table("videos").Select("ID", "User_id", "Play_URL", "Cover_URL", "Favorite_Count", "Comment_Count", "Is_Favorite", "Title", "UNIX_TIMESTAMP(Create_date)").Order("create_date desc").Limit(30).Rows()

	if err != nil {
		fmt.Println("select db failed in func: Feed, err:", err)
		return
	}

	nextTime := time.Now().Unix()

	for rows.Next() {
		var (
			ID             int64
			User_id        int64
			Play_URL       string
			Cover_URL      string
			Favorite_Count int64
			Comment_Count  int64
			Is_Favorite    bool
			Title          string
			Create_date    int64
		)

		err := rows.Scan(&ID, &User_id, &Play_URL, &Cover_URL, &Favorite_Count, &Comment_Count, &Is_Favorite, &Title, &Create_date)
		fmt.Println(Create_date)
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
			continue
		}

		var user User
		result := db.Where("ID = ?", User_id).Take(&user)
		if result.RowsAffected == 0 {
			fmt.Println("In func: get_Feed, Video found no user")
			continue
		}

		v := Video{
			ID:             ID,
			Author:         user,
			Play_URL:       Play_URL,
			Cover_URL:      Cover_URL,
			Favorite_Count: Favorite_Count,
			Comment_Count:  Comment_Count,
			Is_Favorite:    Is_Favorite,
			Title:          Title,
		}
		nextTime = min(nextTime, Create_date)
		videoList = append(videoList, v)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "Feed OK!"},
		VideoList: videoList,
		NextTime:  nextTime,
	})
}

func Get_Feed(r *gin.Engine) {
	r.GET("/douyin/feed/", get_Feed)
}
