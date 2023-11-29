package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"wow-admin/utils/result"
)

// 参数合法性校验
func Validate(c *gin.Context, data any) {
	validMsg := Validator.Validate(data)
	if validMsg != "" {
		result.ReturnJSON(c, http.StatusOK, result.ERROR_INVALID_PARAM, validMsg, nil)
		panic(nil)
	}
}

// Json 绑定
func BindJson[T any](c *gin.Context) (data T) {
	if err := c.ShouldBindJSON(&data); err != nil {
		Logger.Error("ShouldBindJSON", zap.Error(err))
		panic(result.ERROR_REQUEST_PARAM)
	}
	return
}

// Json 绑定验证 + 合法性校验
func BindValidJson[T any](c *gin.Context) (data T) {
	if err := c.ShouldBindJSON(&data); err != nil {
		Logger.Error("ShouldBindJSON", zap.Error(err))
		panic(result.ERROR_REQUEST_PARAM)
	}
	Validate(c, &data)
	return data
}
