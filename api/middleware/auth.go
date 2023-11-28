package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"wow-admin/utils"
	"wow-admin/utils/result"
)

func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.Logger.Info(c.Request.URL.Path)
		requestPath := c.Request.URL.Path
		needLogin := ignoreRoute(requestPath)

		if needLogin {
			c.Next()
			return
		}

		// Authorization: Bearer xxx
		token := c.Request.Header.Get("Authorization")

		// token 为空
		if token == "" {
			result.SendCode(c, result.ERROR_TOKEN_NOT_EXIST)
			c.Abort()
			return
		}

		// token 的正确格式: `Bearer [tokenString]`
		parts := strings.Split(token, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.SendCode(c, result.ERROR_TOKEN_TYPE_WRONG)
			c.Abort()
			return
		}

		// parts[1] 是获取到的 tokenString, 使用 JWT 解析函数解析它
		claims, err := utils.GetJWT().ParseToken(parts[1])
		// token 解析失败
		if err != nil {
			result.SendData(c, result.ERROR_TOKEN_WRONG, err.Error())
			c.Abort()
			return
		}

		// 判断 token 已过期
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			result.SendCode(c, result.ERROR_TOKEN_RUNTIME)
			c.Abort()
			return
		}

		// 将当前请求的相关信息保存到请求的上下文 c 上
		// 后续的处理函数可以用过 c.Get("xxx") 来获取当前请求的用户信息
		c.Set("user_info_id", claims.UserId)
		c.Set("role", claims.Role)
		c.Set("uuid", claims.UUID)

		c.Next()
	}
}

/*
*
白名单路由
*/
func ignoreRoute(requestPath string) bool {
	if strings.HasPrefix(requestPath, "/v1/common") {
		return true
	}
	return false
}
