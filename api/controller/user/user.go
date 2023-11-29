package user

import (
	"github.com/gin-gonic/gin"
	"wow-admin/api/router"
	"wow-admin/model/dto"
	"wow-admin/service"
	"wow-admin/utils"
	"wow-admin/utils/result"
)

type UserController struct {
	userService service.UserService
}

func init() {
	u := &UserController{}
	base := router.GetWebRouterGroup()
	base.POST("/login", u.login)
	base.POST("/register", u.register)

	userRouter := base.Group("/user")
	userRouter.GET("getUserInfo", u.getUserInfo)
}

// 后台登录
func (c *UserController) login(ctx *gin.Context) {
	loginReq := utils.BindValidJson[dto.Login](ctx)

	loginVO, code := c.userService.Login(ctx, loginReq.Username, loginReq.Password)

	result.SendData(ctx, code, loginVO)
}

func (c *UserController) register(ctx *gin.Context) {
	registerReq := utils.BindValidJson[dto.Register](ctx)
	code := c.userService.Register(registerReq)
	result.SendCode(ctx, code)
}

func (c *UserController) getUserInfo(ctx *gin.Context) {
	username := "Test"
	password := "123"

	user := make(map[string]string)
	user["username"] = username
	user["password"] = password

	result.SendData(ctx, 0, user)
}
