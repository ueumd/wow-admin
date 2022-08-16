package common

import (
	"github.com/gin-gonic/gin"
	"time"
	"wow-admin/api/router"
	"wow-admin/constant"
	"wow-admin/global/response"
	"wow-admin/service"
)


type CommonController struct {
	response.ResponseUtil
	commService  service.CommService
}

func init()  {
	c := &CommonController{}
	commGroup := router.GetWebRouter().Group("/v1/common")
	commGroup.GET("/getSystemDate", c.GetSystemDate)
	commGroup.GET("/getDictionaryData", c.getDictionaryData)
	commGroup.GET("/getAllSupportCityList", c.GetAllSupportCityList)
}

func (c* CommonController) GetSystemDate(ctx *gin.Context) {
	timestamp := time.Now().Unix()
	formatTimeStr:=time.Unix(timestamp,0).Format("2006-01-02 15:04:05")
	result := make(map[string]interface{})
	result["timestamp"] = timestamp * 1000
	result["date"] = formatTimeStr
	c.Json200OK(ctx, result)
}


func (c * CommonController) getDictionaryData(ctx *gin.Context) {
	result := make(map[string]interface{})
	result["language"] = constant.Language.RangeIntKeyValue()
	c.Json200OK(ctx, result)
}


func (c *CommonController) GetAllSupportCityList(ctx *gin.Context) {
	resp, _ := c.commService.GetAllSupportCityList()
	c.Json200OK(ctx, resp)
}
