package controller

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

// 对应数据库中的User表
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
	// Videos          []Video   `gorm:"ForeignKey:DeviceHexCode;AssociationForeignKey:DeviceHexCode" json:"video"`
}

// deprecated
type User_From_Pamphlet struct {
	ID              int64  `json:"id,omitempty"`               // 用户id
	Avatar          string `json:"avatar,omitempty"`           // 用户头像
	BackgroundImage string `json:"background_image,omitempty"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count,omitempty"`   // 喜欢数
	FollowCount     int64  `json:"follow_count,omitempty"`     // 关注总数
	FollowerCount   int64  `json:"follower_count,omitempty"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow,omitempty"`        // true-已关注，false-未关注
	Name            string `json:"name,omitempty"`             // 用户名称
	Signature       string `json:"signature,omitempty"`        // 个人简介
	TotalFavorited  int64  `json:"total_favorited,omitempty"`  // 获赞数量
	WorkCount       int64  `json:"work_count,omitempty"`       // 作品数
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

// 对应数据库中的视频表
type Video struct {
	ID             int64     `json:"id"`                               // 视频唯一标识
	User_id        int64     `json:"user_id"`                          // 作者id
	User_token     string    `json:"user_token"`                       // 作者token
	User_name      string    `json:"User_name"`                        // 作者名
	Author         User      `gorm:"foreignKey:User_id" json:"author"` // 视频作者信息
	Comment_Count  int64     `json:"comment_count"`                    // 视频的评论总数
	Cover_URL      string    `json:"cover_url"`                        // 视频封面地址
	Favorite_Count int64     `json:"favorite_count"`                   // 视频的点赞总数
	Is_Favorite    bool      `json:"is_favorite"`                      // true-已点赞，false-未点赞
	Play_URL       string    `json:"play_url"`                         // 视频播放地址
	Title          string    `json:"title"`                            // 视频标题
	Create_Date    time.Time `json:"Create_Date"`                      // 记录创建日期
}

// 对应数据库中的点赞表
type Favorite struct {
	ID       int64 `json:"id"`       // 唯一标识
	User_id  int64 `json:"user_id"`  // 作者id
	Video_id int64 `json:"video_id"` // 此作者点赞的视频id
}

// 对应数据库中的视频评论表
type Comment struct {
	ID          int64     `json:"id"`          // 唯一标识
	User_id     int64     `json:"user_id"`     // 作者id
	Video_id    int64     `json:"video_id"`    // 视频id
	Content     string    `json:"content"`     // 评论内容
	Create_Date time.Time `json:"create_date"` // 记录创建日期
	Is_delete   bool      `json:"is_delete"`   // 逻辑删除
}

// 返回前端
type CommentList struct {
	Id          int64  `json:"id,omitempty"`
	User        User   `json:"user,omitempty"`
	Content     string `json:"content,omitempty"`
	Create_Date string `json:"create_date,omitempty"`
}

// 返回前端
type CommentListResponse struct {
	Response
	CommentList []CommentList `json:"comment_list,omitempty"`
}

// 返回前端
type CommentActionResponse struct {
	Response
	CommentList CommentList `json:"comment,omitempty"`
}

var DemoUser = User{
	ID:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
