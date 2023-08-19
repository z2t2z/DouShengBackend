package main

import (
	"time"
)

// 通用返回响应
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// 用户注册、登录
type UserLoginResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

type UserInfo struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	User       User   `json:"user"`        // 用户信息
}

// User Info
type User struct {
	ID              int64     `json:"id"`               // 用户id
	Avatar          string    `json:"avatar"`           // 用户头像
	BackgroundImage string    `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64     `json:"favorite_count"`   // 喜欢数
	FollowCount     int64     `json:"follow_count"`     // 关注总数
	FollowerCount   int64     `json:"follower_count"`   // 粉丝总数
	IsFollow        bool      `json:"is_follow"`        // true-已关注，false-未关注
	Name            string    `json:"name"`             // 用户名称
	Password        string    `json:"password"`         // 密码
	Token           string    `json:"token"`            // name+password
	Signature       string    `json:"signature"`        // 个人简介
	TotalFavorited  int64     `json:"total_favorited"`  // 获赞数量
	WorkCount       int64     `json:"work_count"`       // 作品数
	CreateDate      time.Time `json:"Create_Date"`      // 记录创建日期
}


type ApifoxModel struct {
	NextTime   *int64  `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Video
type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

// var DemoVideos = []Video{
// 	{
// 		Id:            1,
// 		Author:        DemoUser,
// 		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
// 		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
// 		FavoriteCount: 0,
// 		CommentCount:  0,
// 		IsFavorite:    false,
// 	},
// }

// 视频作者信息
//
// User
// type User struct {
// 	Avatar          string `json:"avatar"`          // 用户头像
// 	BackgroundImage string `json:"background_image"`// 用户个人页顶部大图
// 	FavoriteCount   int64  `json:"favorite_count"`  // 喜欢数
// 	FollowCount     int64  `json:"follow_count"`    // 关注总数
// 	FollowerCount   int64  `json:"follower_count"`  // 粉丝总数
// 	ID              int64  `json:"id"`              // 用户id
// 	IsFollow        bool   `json:"is_follow"`       // true-已关注，false-未关注
// 	Name            string `json:"name"`            // 用户名称
// 	Signature       string `json:"signature"`       // 个人简介
// 	TotalFavorited  string `json:"total_favorited"` // 获赞数量
// 	WorkCount       int64  `json:"work_count"`      // 作品数
// }
