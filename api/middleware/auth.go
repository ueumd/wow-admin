package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ueumd/logger"
	"net/http"
	"strings"
	"wow-admin/global/cerrors"
	"wow-admin/service"
)

func AuthLogin() gin.HandlerFunc  {
	return func(context *gin.Context) {
		logger.Infoln(context.Request.URL.Path)
		requestPath := context.Request.URL.Path
		needLogin := isNeedLogin(requestPath)

		token := context.Request.Header.Get("X-TOKEN")

		if token == "" && needLogin {
			context.JSONP(http.StatusUnauthorized, gin.H{
				"code": cerrors.ErrDefaultNeedLogin,
				"message": "您还未登陆请先登陆",
			})
			context.Abort()
			return
		}

		// 解析token
		claim, err := service.DecodeToken(token)
		if err != nil && needLogin {
			logger.ErrorF("Token auth error: v%", err)
			context.JSONP(http.StatusOK, gin.H{
				"code": cerrors.ErrDefaultToken,
				"message": "请重新登陆！",
			})
		}


		loginService := &service.LoginService{}
		if claim != nil {
			tk, err := loginService.GetLoginToken(claim.UserId)
			if (err != nil || tk != token) && needLogin {
				context.JSON(http.StatusOK, gin.H{"code": cerrors.ErrDefaultToken, "msg": "请重新登陆！"})
				context.Abort()
				return
			}

			context.Set("Token", claim)
			context.Set("token", token)
			context.Set("userId", claim.UserId)
			context.Set("unionId", claim.UnionId)
			context.Set("phone", claim.Phone)
		}


		context.Next()
	}
}

func isNeedLogin(requestPath string) bool {
	if strings.HasPrefix(requestPath, "/v1/user") {
		return true
	}

	if strings.HasPrefix(requestPath, "/v1/collects") {
		return true
	}

	return false
}