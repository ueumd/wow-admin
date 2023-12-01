package user

import (
	"github.com/gin-gonic/gin"
	"wow-admin/api/router"
	"wow-admin/model/req"
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
	userRouter.GET("logout", u.logout)
}

// 后台登录
func (c *UserController) login(ctx *gin.Context) {
	loginReq := utils.BindValidJson[req.Login](ctx)

	loginVO, code := c.userService.Login(ctx, loginReq.Username, loginReq.Password)

	result.SendData(ctx, code, loginVO)
}

func (c *UserController) logout(ctx *gin.Context) {
	c.userService.Logout(ctx)
	result.Success(ctx)
}

// 注册
func (c *UserController) register(ctx *gin.Context) {
	// data := utils.BindValidJson[dto.Register](ctx)
	// validMsg := utils.Validate(ctx, &data)
	// if validMsg != "" {
	//	result.ReturnJSON(ctx, http.StatusOK, result.ERROR_INVALID_PARAM, validMsg, nil)
	//	return
	// }
	// result.SendCode(ctx, c.userService.Register(data))

	result.SendCode(ctx, c.userService.Register(utils.BindValidJson[req.Register](ctx)))
}

func (c *UserController) getUserInfo(ctx *gin.Context) {
	result.SuccessData(ctx, c.userService.GetUserInfo(utils.GetFromContext[int](ctx, "userId")))
}
