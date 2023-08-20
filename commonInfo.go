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
	ID              int64  `json:"id"`               // 用户id
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Password        string `json:"password"`         // 密码
	Token           string `json:"token"`            // name+password
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
	// Videos          []Video   `gorm:"ForeignKey:DeviceHexCode;AssociationForeignKey:DeviceHexCode" json:"video"`
	CreateDate time.Time `json:"Create_Date"` // 记录创建日期
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

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Video
type Video struct {
	ID             int64     `json:"id"`                               // 视频唯一标识
	User_id        int64     `json:"user_id"`                          // 作者id
	User_token     string    `json:"user_token"`                       // 作者token
	Author         User      `gorm:"foreignKey:User_id" json:"Author"` // 视频作者信息
	Comment_Count  int64     `json:"comment_count"`                    // 视频的评论总数
	Cover_URL      string    `json:"cover_url"`                        // 视频封面地址
	Favorite_Count int64     `json:"favorite_count"`                   // 视频的点赞总数
	Is_Favorite    bool      `json:"is_favorite"`                      // true-已点赞，false-未点赞
	Play_URL       string    `json:"play_url"`                         // 视频播放地址
	Title          string    `json:"title"`                            // 视频标题
	Create_Date    time.Time `json:"Create_Date"`                      // 记录创建日期
}

// 点赞表
type Favorite struct {
	ID       int64 `json:"id"`       // 唯一标识
	User_id  int64 `json:"user_id"`  // 作者id
	Video_id int64 `json:"video_id"` // 此作者点赞的视频id
}

// 视频评论表
type Comment struct {
	ID          int64     `json:"id"`          // 唯一标识
	User_id     int64     `json:"user_id"`     // 作者id
	Video_id    int64     `json:"video_id"`    // 视频id
	Content     string    `json:"content"`     // 评论内容
	Create_Date time.Time `json:"Create_Date"` // 记录创建日期
	Is_delete   bool      `json:"is_delete"`   // 逻辑删除
}

// type Video struct {
// 	Author        User   `json:"author"`         // 视频作者信息
// 	Comment_Count  int64  `json:"comment_count"`  // 视频的评论总数
// 	Cover_URL      string `json:"cover_url"`      // 视频封面地址
// 	Favorite_Count int64  `json:"favorite_count"` // 视频的点赞总数
// 	ID            int64  `json:"id"`             // 视频唯一标识
// 	Is_Favorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
// 	Play_URL       string `json:"play_url"`       // 视频播放地址
// 	Title         string `json:"title"`          // 视频标题
// }

var DemoUser = User{
	ID:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
