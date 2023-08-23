package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// util "myproj/controller/utility"
)

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	var user User
	var exist bool
	if user, exist = check_User_Login_Status(token); !exist {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}

	// to_user_sid := c.Query("to_user_id")
	// action_type := c.Query("action_type")

	// if action_type == "1" {

	// } else if action_type == "2" {

	// }
	fmt.Println(user.ID)

}
