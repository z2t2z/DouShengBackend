package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 点赞：Favorite表新增一条user<->video记录、video对应的点赞数+1，user对应的点赞数+1
// 可以写成事务
func Like(c *gin.Context) {
	token := c.Query("token")
	user, has_user := getUser(token)

	if !has_user {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}

	video_sid := c.Query("video_id")
	action_type := c.Query("action_type")

	// fmt.Println(video_sid, " ", action_type)

	video_id, _ := strconv.ParseInt(video_sid, 10, 64)
	user_id := user.ID
	favorite := Favorite{
		User_id:  user_id,
		Video_id: video_id,
	}

	fmt.Println(video_id, " ", user_id)

	if action_type == "1" { // 点赞
		result := db.Where(Favorite{
			User_id:  user_id,
			Video_id: video_id,
		}).Take(&favorite)

		if result.RowsAffected == 1 {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "You have liked this video already!",
			})
			return
		}

		if err := db.Create(&favorite).Error; err != nil {
			fmt.Println("Error In function: Register, Insert failed!\n", err)
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Like error 1",
			})
			return
		}

		// 更新user表点赞+1
		if err := db.Model(&user).Where("token = ? ", token).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Like error 1",
			})
			return
		}

		video := Video{
			ID: video_id,
		}
		// 更新video表点赞+1
		if err := db.Model(&video).Where("ID = ? ", video_id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Like error 2",
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "Like success!",
		})

	} else if action_type == "2" { // 取消点赞

		// 硬删除：直接在favorite表中抹去记录
		if err := db.Where(Favorite{User_id: user_id, Video_id: video_id}).Delete(&favorite).Error; err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Dislike error!",
			})
		}

		// 更新user表点赞-1
		if err := db.Model(&user).Where("token = ? ", token).Update("favorite_count", gorm.Expr("favorite_count + ?", -1)).Error; err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Dislike error 1",
			})
			return
		}

		video := Video{
			ID: video_id,
		}
		// 更新video表点赞-1
		if err := db.Model(&video).Where("ID = ? ", video_id).Update("favorite_count", gorm.Expr("favorite_count + ?", -1)).Error; err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Dislike error 2",
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "Dislike success!",
		})
	}
}

func User_Like(r *gin.Engine) {
	r.POST("/douyin/favorite/action/", Like)
}
