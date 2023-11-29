package common

import (
	"github.com/gin-gonic/gin"
	"time"
	"wow-admin/api/router"
	"wow-admin/constant"
	"wow-admin/service"
	"wow-admin/utils/result"
)

type CommonController struct {
	menuService service.MenuService
	cityService service.CityService
}

func init() {
	c := &CommonController{}
	r := router.GetWebRouterGroup()
	commGroup := r.Group("/common")
	commGroup.GET("/getSystemDate", c.GetSystemDate)
	commGroup.GET("/getDictionaryData", c.getDictionaryData)
	commGroup.GET("/getAllSupportCityList", c.GetAllSupportCityList)
	commGroup.GET("/menu", c.GetMenu)
}

func (c *CommonController) GetSystemDate(ctx *gin.Context) {
	timestamp := time.Now().Unix()
	formatTimeStr := time.Unix(timestamp, 0).Format(time.DateTime)
	res := make(map[string]interface{})
	res["timestamp"] = timestamp * 1000
	res["date"] = formatTimeStr
	result.SendData(ctx, 0, res)
}

func (c *CommonController) getDictionaryData(ctx *gin.Context) {
	res := make(map[string]interface{})
	res["language"] = constant.Language.RangeIntKeyValue()
	result.SendData(ctx, 0, res)
}

func (c *CommonController) GetAllSupportCityList(ctx *gin.Context) {
	resp, _ := c.cityService.GetAllSupportCityList()
	result.SendData(ctx, 0, resp)
}

func (c *CommonController) GetMenu(ctx *gin.Context) {
	resp := c.menuService.GetTreeList()
	result.SendData(ctx, 0, resp)
}
