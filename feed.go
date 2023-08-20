package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// var DemoVideos = []Video{
// 	{
// 		ID:             1,
// 		Author:         DemoUser,
// 		Play_URL:       "http://192.168.0.105:9090/public/videos/test1.mp4",
// 		Cover_URL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
// 		Favorite_Count: 0,
// 		Comment_Count:  0,
// 		Is_Favorite:    false,
// 		Title:          "none",
// 	},
// }

func Feed(c *gin.Context) {
	// lastest_time := c.Query("latest_time")
	// token = c.Query("token")
	// fmt.Println(lastest_time)

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

// func Feed(c *gin.Context) {
// 	lastest_time := c.Query("latest_time")
// 	token := c.Query("token")
// 	fmt.Println(token, lastest_time)

// 	var DemoUser = User{
// 		ID:            1,
// 		Name:          "TestUser",
// 		FollowCount:   0,
// 		FollowerCount: 0,
// 		IsFollow:      false,
// 	}

// 	var DemoVideos = []Video{
// 		{
// 			ID:             1,
// 			Author:         DemoUser,
// 			Play_URL:       "http://192.168.0.105:9090/public/videos/test1.mp4",
// 			Cover_URL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
// 			Favorite_Count: 0,
// 			Comment_Count:  0,
// 			Is_Favorite:    false,
// 			Title:          "none",
// 		},
// 	}

// 	c.JSON(http.StatusOK, FeedResponse{
// 		Response:  Response{StatusCode: 0, StatusMsg: "OK!"},
// 		VideoList: DemoVideos,
// 		// NextTime:  time.Now().Unix(),
// 	})
// }

func getFeed(r *gin.Engine) {
	r.GET("/douyin/feed/", Feed)
}
