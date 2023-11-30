package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"wow-admin/utils/result"
)

/**
gin 上下文
*/

// 参数合法性校验
func Validate(c *gin.Context, data any) {
	validMsg := Validator.Validate(data)
	if validMsg != "" {
		result.ReturnJSON(c, http.StatusOK, result.ERROR_INVALID_PARAM, validMsg, nil)

		// go 1.21版本中 panic(nil) 会得到 panic called with nil argument
		// src/runtime/panic.go 定义了一个名为PanicNilError的新Error
		// func (*PanicNilError) Error() string { return "panic called with nil argument" }
		// func (*PanicNilError) RuntimeError() {}
		// panic(nil)

		panic("wow oh no!")
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
	// Json 绑定
	if err := c.ShouldBindJSON(&data); err != nil {
		Logger.Error("BindValidJson", zap.Error(err))
		panic(result.ERROR_REQUEST_PARAM)
	}
	Validate(c, &data)
	return data
}

// 从 Gin Context 上获取值, 该值是 JWT middleware 解析 Token 后设置的
// 如果该值不存在, 说明 Token 有问题
func GetFromContext[T any](c *gin.Context, key string) T {
	val, exist := c.Get(key)
	if !exist {
		panic(result.ERROR_TOKEN_RUNTIME)
	}
	return val.(T)
}

// 从 Context 获取 Int 类型 Param 参数
func GetIntParam(c *gin.Context, key string) int {
	val, err := strconv.Atoi(c.Param(key))
	if err != nil {
		Logger.Error("GetIntParam", zap.Error(err))
		panic(result.ERROR_REQUEST_PARAM)
	}
	return val
}
