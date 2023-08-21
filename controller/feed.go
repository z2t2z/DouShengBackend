package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 这里没有限制返回的具体时间戳
func get_Feed(c *gin.Context) {
	var videoList []Video

	rows, err := db.Table("videos").Select("ID", "Play_URL", "Cover_URL", "Favorite_Count", "Comment_Count", "Is_Favorite", "Title").Order("create_date desc").Limit(30).Rows()

	if err != nil {
		fmt.Println("select db failed in func: Feed, err:", err)
		return
	}

	for rows.Next() {
		var (
			ID             int64
			Play_URL       string
			Cover_URL      string
			Favorite_Count int64
			Comment_Count  int64
			Is_Favorite    bool
			Title          string
		)

		err := rows.Scan(&ID, &Play_URL, &Cover_URL, &Favorite_Count, &Comment_Count, &Is_Favorite, &Title)
		fmt.Println(Play_URL)
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
			continue
		}

		v := Video{
			ID:             ID,
			Play_URL:       Play_URL,
			Cover_URL:      Cover_URL,
			Favorite_Count: Favorite_Count,
			Comment_Count:  Comment_Count,
			Is_Favorite:    Is_Favorite,
			Title:          Title,
		}

		videoList = append(videoList, v)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "Feed OK!"},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}

func Get_Feed(r *gin.Engine) {
	r.GET("/douyin/feed/", get_Feed)
}
