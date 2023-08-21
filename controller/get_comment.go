package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func get_Comment(c *gin.Context) {
	token := c.Query("token")
	video_sid := c.Query("video_id")

	_, login := check_User_Login_Status(token)
	if !login {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist!",
		})
		return
	}

	video_id, _ := strconv.ParseInt(video_sid, 10, 64)
	fmt.Println(token, video_id)

	rows, err := db.Table("comments").Select("video_id", "user_id", "content", "create_date").Where("video_id = ?", video_id).Order("create_date desc").Limit(30).Rows()
	if err != nil {
		fmt.Println("select db failed in func: get_Comment, err:", err)
		return
	}

	var commentList []CommentList
	for rows.Next() {
		var (
			id          int64
			user_id     int64
			content     string
			create_date time.Time
		)

		err := rows.Scan(&id, &user_id, &content, &create_date)
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
			continue
		}

		var user User

		result := db.Where("token = ?", token).Take(&user)
		if result.RowsAffected == 0 {
			fmt.Println("In func: get_Comment, found no user")
			continue
		}

		commentList = append(commentList, CommentList{
			id,
			user, 
			content,
			create_date.String(),
		})
		fmt.Println(content)
	}
	rows.Close()

	c.JSON(http.StatusOK, CommentListResponse{
		Response: Response{StatusCode: 0},
		CommentList: commentList,
	})

}

func Get_Comment(r *gin.Engine) {
	r.GET("/douyin/comment/list/", get_Comment)
}
