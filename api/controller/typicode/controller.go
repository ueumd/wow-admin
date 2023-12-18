package typicode

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"wow-admin/api/router"
	"wow-admin/utils"
	"wow-admin/utils/result"
)

type TypicodeController struct {
}

type articleBase struct {
	UserId int    `json:"userId" validate:"required" label:"用户ID"`
	Title  string `json:"title" validate:"required" label:"文章标题"`
	Body   string `json:"body" validate:"required" label:"文章内容"`
}

type articleInfo struct {
	Id int `json:"id"`
	articleBase
}

func init() {
	c := &TypicodeController{}
	r := router.GetWebRouterGroup()
	group := r.Group("/typicode")

	group.GET("/article", c.list)
	group.POST("/article", c.add)
	group.PUT("/article", c.put)
}

func (c *TypicodeController) list(ctx *gin.Context) {
	var list []articleInfo
	response, err := resty.New().R().SetResult(&list).Get("http://jsonplaceholder.typicode.com/posts")
	if err != nil {
		result.Send500Error(ctx, "服务异常")
		utils.Logger.Error("服务异常: ", zap.Error(err))
		return
	}
	fmt.Println("Status: ", response.Status())
	result.SendData(ctx, 0, list)
}

func (c *TypicodeController) add(ctx *gin.Context) {
	var article articleBase
	articleReq := utils.BindValidJson[articleBase](ctx)

	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(&articleReq).
		SetResult(&article).
		Post("http://jsonplaceholder.typicode.com/posts")
	if err != nil {
		result.Send500Error(ctx, "服务异常")
		utils.Logger.Error("服务异常: ", zap.Error(err))
		return
	}
	fmt.Println("Status: ", response.Status())
	result.SendData(ctx, 0, &article)
}

func (c *TypicodeController) put(ctx *gin.Context) {
	var article articleInfo
	articleReq := utils.BindValidJson[articleInfo](ctx)

	url := fmt.Sprintf("http://jsonplaceholder.typicode.com/posts/%d", articleReq.Id)
	fmt.Println("url: ", url)

	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(&articleReq).
		SetResult(&article).
		Post(url)

	if err != nil {
		result.Send500Error(ctx, "服务异常")
		utils.Logger.Error("服务异常: ", zap.Error(err))
		return
	}
	fmt.Println("Status: ", response.Status())
	result.SendData(ctx, 0, &article)
}
