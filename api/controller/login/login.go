package login

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"wow-admin/api/router"
	"wow-admin/global/response"
	"wow-admin/model/loginreq"
	"wow-admin/service"
	"wow-admin/utils"
)

type LoginController struct {
	loginService service.LoginService
	response.ResponseUtil
}

func init() {
	u := &LoginController{}
	loginGroup := router.GetWebRouter().Group("v1/login")
	loginGroup.POST("phoneLoginTest", u.PhoneLoginTest)
	loginGroup.POST("phoneLogin", u.PhoneLogin)
}

func (c *LoginController) PhoneLoginTest(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	user := make(map[string]string)
	user["username"] = username
	user["password"] = password

	c.Json200OK(ctx, user)
}

func (u *LoginController) PhoneLogin(ctx *gin.Context) {
	phoneReq := &loginreq.PhoneLoginRequest{}
	err := ctx.BindJSON(phoneReq)
	if err != nil {
		u.Json400Error(ctx, "出现意外，填写格式问题")
		return
	}

	if m, err := regexp.Match(utils.PhoneRegexp, []byte(phoneReq.Phone)); !m || err != nil {
		u.Json400Error(ctx, "请填写正确的手机号")
		return
	}

	//if m, err := regexp.Match(utils.PasswordRegexp, []byte(phoneReq.Password)); (!m || err != nil) && len(phoneReq.CheckCode) < 5 {
	//	u.Json400Error(ctx, "请填写正确密码或者验证码")
	//	return
	//}

	token, err := u.loginService.PhoneLogin(phoneReq)
	if err != nil {
		u.JsonErrorResult(ctx, err)
	} else {
		u.Json200OK(ctx, token)
	}
}
