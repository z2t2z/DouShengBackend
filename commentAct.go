package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Post_Comment(c *gin.Context) {
	token := c.Query("token")

	user, has_user := getUser(token)
	if !has_user {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist!",
		})
		return
	}

	video_sid := c.Query("video_id")
	action_type := c.Query("action_type")

	video_id, _ := strconv.ParseInt(video_sid, 10, 64)

	fmt.Println(user.ID, token, video_id)

	if action_type == "1" { // 新增评论
		comment_text := c.Query("comment_text")
		fmt.Println(comment_text)

		comment := Comment {
			User_id: user.ID, 
			Video_id: video_id,
			Content: comment_text,
		}

		if err := db.Select("User_id", "Video_id", "Content").Create(&comment).Error; err != nil {
			fmt.Println("Error In function: Post_Comment, Insert failed!\n", err)
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Post_Comment error 1",
			})
			return
		}
		// video表评论数+1

		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "Comment success!",
		})
	} else if action_type == "2" { //删除评论



		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "Delete a comment success!",
		})
	}
}

func Comment_Action(r *gin.Engine) {
	r.POST("/douyin/comment/action/", Post_Comment)
}
