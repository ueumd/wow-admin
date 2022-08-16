package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ueumd/logger"
	"strings"
	"wow-admin/global/cerrors"
	"wow-admin/utils/cast"
)

type ResponseUtil struct {
}

func (res *ResponseUtil) ErrorLog(ctx *gin.Context, err error) {
	if err == nil {
		return
	}
	errorStrings := err.Error()
	strValues := strings.Split(errorStrings, "\n")
	logEntry := logger.WithFiled("lrequest", ctx.Request.URL.Path)
	for _, v := range strValues {
		logEntry.Errorln(strings.Replace(v, "\t", "", 1))
	}
}

type CommResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res *ResponseUtil) JsonResult(ctx *gin.Context, code int, message string, data interface{}) {
	if cast.IsNil(data) {
		data = make(map[string]interface{})
	}

	ctx.JSON(200, &CommResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func (res *ResponseUtil) Json500Error(ctx *gin.Context, message string) {
	ctx.JSON(200, &CommResponse{
		Code:    500,
		Message: message,
		Data:    make(map[string]interface{}),
	})
}

func (res *ResponseUtil) Json400Error(ctx *gin.Context, message string) {
	if message == "" {
		message = "出现意外，填写格式问题"
	}
	ctx.JSON(200, &CommResponse{
		Code:    400,
		Message: message,
		Data:    make(map[string]interface{}),
	})
}

func (res *ResponseUtil) Json200OK(ctx *gin.Context, data interface{}) {
	if cast.IsNil(data) {
		data = make(map[string]interface{})
	}

	ctx.JSON(200, &CommResponse{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func (res *ResponseUtil) JsonErrorResult(ctx *gin.Context, err error) {
	res.ErrorLog(ctx, err) //print error log
	var serviceErr cerrors.ServiceError
	if errors.As(err, &serviceErr) {
		ctx.JSON(200, &CommResponse{
			Code:    serviceErr.ErrCode(),
			Message: serviceErr.ErrMessage(),
			Data:    serviceErr.ErrData(),
		})
	} else {
		ctx.JSON(200, &CommResponse{
			Code:    500,
			Message: "服务器内部错误",
			Data:    make(map[string]interface{}),
		})
	}
}
