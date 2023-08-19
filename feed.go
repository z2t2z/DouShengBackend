package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

var DemoUser = User{
	ID:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

// file:///E:/ByteDanceSummerCamp/Data/videos/0_test1.mp4
// \\192.168.0.105\public
//	http://localhost:8888/videos/test1.mp4
// https://blog.csdn.net/weixin_41010198/article/details/88055078
// 基于Gin框架的极简版抖音开发
// 基于Gin框架的极简版抖音开发
// 基于Gin框架的极简版抖音开发
// 基于Gin框架的极简版抖音开发

var DemoVideos = []Video{
	{
		ID:            1,
		Author:        DemoUser,
		PlayURL:       "http://192.168.0.105:9090/public/test2.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "none",
	},
}

func Feed(c *gin.Context) {
	// lastest_time := c.Query("latest_time")
	// token = c.Query("token")
	// fmt.Println(lastest_time)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "OK!"},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}

func getFeed(r *gin.Engine) {
	r.GET("/douyin/feed/", Feed)
}
