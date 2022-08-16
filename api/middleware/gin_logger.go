package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ueumd/logger"
	"time"
)

// 路由日志中间件
func RouterLogger() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		var statusColor, methodColor, resetColor string
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()
		if raw != "" {
			path = path + "?" + raw
		}

		logger.InfoF("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %s %s",
			end.Format("2006/01/02 - 15:04:05"),
			statusColor, statusCode, resetColor,
			latency,
			clientIP,
			methodColor, method, resetColor,
			path,
			comment,
		)
	}
}
