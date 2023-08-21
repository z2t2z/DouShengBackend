package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get_Favorite_List(c *gin.Context) {
	user_sid := c.Query("user_id")
	token := c.Query("token")

	user_id, _ := strconv.ParseInt(user_sid, 10, 64)

	fmt.Println(user_id, token)

	/* 
		unimplemented
	*/

}

func Get_Favorite_List(r *gin.Engine) {
	r.GET("/douyin/favorite/list/", get_Favorite_List)
}
