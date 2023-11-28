package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ReturnJSON(c *gin.Context, httpCode, code int, message string, data any) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Send(c *gin.Context, httpCode, code int, data any) {
	ReturnJSON(c, httpCode, code, GetMsg(code), data)
}

func SendCode(c *gin.Context, code int) {
	Send(c, http.StatusOK, code, nil)
}

func SendData(c *gin.Context, code int, data any) {
	Send(c, http.StatusOK, code, data)
}

func SuccessData(c *gin.Context, data any) {
	Send(c, http.StatusOK, OK, data)
}

func Success(c *gin.Context) {
	Send(c, http.StatusOK, OK, nil)
}

func Send400Error(c *gin.Context, message string) {
	if message == "" {
		message = "出现意外，填写格式问题"
	}
	ReturnJSON(c, http.StatusOK, 400, message, nil)
}

func Send500Error(c *gin.Context, message string) {
	ReturnJSON(c, http.StatusOK, 500, message, make(map[string]interface{}))
}
