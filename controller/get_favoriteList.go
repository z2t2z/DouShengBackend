package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get_Favorite_List(c *gin.Context) {
	user_sid := c.Query("user_id")
	token := c.Query("token")

	_, login := check_User_Login_Status(token)
	if !login {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist!",
		})
		return
	}

	user_id, _ := strconv.ParseInt(user_sid, 10, 64)
	rows, err := db.Table("favorites").Select("video_id").Where("user_id = ?", user_id).Rows()

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Err in func: get_Favorite_List!",
		})
		return
	}

	var videoList []Video
	for rows.Next() {
		var video_id int64
		err := rows.Scan(&video_id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var video Video
		result := db.Where("ID = ?", video_id).Take(&video)
		if result.RowsAffected == 0 {
			fmt.Println("In func: get_Favorite_List, found no video")
			continue
		}

		var user User
		result = db.Where("ID = ?", video.User_id).Take(&user)
		if result.RowsAffected == 0 {
			fmt.Println("In func: get_Favorite_List, Video found no user")
			continue
		}

		v := Video{
			ID:             video.ID,
			Author:         user,
			Play_URL:       video.Play_URL,
			Cover_URL:      video.Cover_URL,
			Favorite_Count: video.Favorite_Count,
			Comment_Count:  video.Comment_Count,
			Is_Favorite:    video.Is_Favorite,
			Title:          video.Title,
		}

		videoList = append(videoList, v)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "get_Favorite_List OK!"},
		VideoList: videoList,
	})

}

func Get_Favorite_List(r *gin.Engine) {
	r.GET("/douyin/favorite/list/", get_Favorite_List)
}
