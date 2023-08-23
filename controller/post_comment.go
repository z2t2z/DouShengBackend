package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func post_Comment(c *gin.Context) {
	token := c.Query("token")

	user, login := check_User_Login_Status(token)
	if !login {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist!",
		})
		return
	}

	video_sid := c.Query("video_id")
	action_type := c.Query("action_type")
	video_id, _ := strconv.ParseInt(video_sid, 10, 64)

	if action_type == "1" { // 新增评论
		comment_text := c.Query("comment_text")
		fmt.Println(comment_text)

		comment := Comment{
			User_id:  user.ID,
			Video_id: video_id,
			Content:  comment_text,
		}

		if err := db.Select("User_id", "Video_id", "Content").Create(&comment).Error; err != nil {
			fmt.Println("Error In function: Post_Comment, Insert failed!\n", err)
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Post_Comment error 1",
			})
			return
		}
		fmt.Println("receive comment id is : ", comment.ID)
		// video表评论数+1
		var video Video
		if err := db.Model(&video).Where("id = ? ", video_id).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "post_Comment error 2",
			})
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
			CommentList: CommentList{
				Id:          comment.ID,
				User:        user,
				Content:     comment_text,
				Create_Date: comment.Create_Date.String(),
			}})

	} else if action_type == "2" { //删除评论
		comment_sid := c.Query("comment_id")
		comment_id, _ := strconv.ParseInt(comment_sid, 10, 64)
		fmt.Println("Before delete a comment, receive id is : ", comment_id)

		// comment表删除记录
		var comment Comment
		if err := db.Where(Comment{ID: comment_id}).Delete(&comment).Error; err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "post_Comment error 4",
			})
			return
		}
		// video表评论数-1
		var video Video
		if err := db.Model(&video).Where("id = ? ", video_id).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "post_Comment error 3",
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "Delete a comment success!",
		})
	}
}

func Post_Comment(r *gin.Engine) {
	r.POST("/douyin/comment/action/", post_Comment)
}
