package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"wow-admin/utils"
	"wow-admin/utils/result"
)

// recover 项目可能出现的 panic, 并使用 zap 记录相关日志
func ErrorRecovery(stack bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil && err != "wow oh no!" {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 处理 panic(xxx) 的操作
				// err 类型断言
				if code, ok := err.(int); ok { // panic(code) 根据错误码获取 msg
					result.SendCode(ctx, code)
				} else if msg, ok := err.(string); ok { // panic(string) 返回 string
					result.ReturnJSON(ctx, http.StatusOK, result.FAIL, msg, nil)
				} else if e, ok := err.(error); ok { // panic(error) 发送消息
					result.ReturnJSON(ctx, http.StatusOK, result.FAIL, e.Error(), nil)
				} else { // 其他
					result.Send(ctx, http.StatusOK, result.FAIL, nil)
				}

				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					utils.Logger.Error(ctx.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = ctx.Error(err.(error)) // nolint: err check
					ctx.Abort()
					return
				}

				if stack {
					utils.Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					utils.Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}
