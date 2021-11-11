package user

import (
	"github.com/gin-gonic/gin"
	"wow-admin/api/router"
	"wow-admin/global/response"
)

type UserController struct {
	response.ResponseUtil
}

func init()  {
	u := &UserController{}
	loginGroup := router.GetWebRouter().Group("v1/user")
	loginGroup.GET("getUserInfo", u.getUserInfo)
}

func (c *UserController) getUserInfo(ctx *gin.Context)  {
	username := "Test"
	password := "123"

	user := make(map[string]string)
	user["username"] = username
	user["password"] = password

	c.Json200OK(ctx, user)
}