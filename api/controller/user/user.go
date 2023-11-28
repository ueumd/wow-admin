package user

import (
	"github.com/gin-gonic/gin"
	"wow-admin/api/router"
	"wow-admin/utils/result"
)

type UserController struct {
}

func init() {
	u := &UserController{}
	loginGroup := router.GetWebRouter().Group("v1/user")
	loginGroup.GET("getUserInfo", u.getUserInfo)
}

func (c *UserController) getUserInfo(ctx *gin.Context) {
	username := "Test"
	password := "123"

	user := make(map[string]string)
	user["username"] = username
	user["password"] = password

	result.SendData(ctx, 0, user)
}
