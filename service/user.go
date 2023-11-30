package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"wow-admin/config"
	"wow-admin/dao"
	"wow-admin/model"
	"wow-admin/model/dto"
	"wow-admin/model/req"
	"wow-admin/model/vo"
	"wow-admin/utils"
	"wow-admin/utils/result"
)

type UserService struct {
}

// 后台用户登录
func (s *UserService) Login(c *gin.Context, username string, password string) (loginVO vo.LoginVO, code int) {
	userAuth := dao.GetOne(model.UserAuth{}, "username", username)
	if userAuth.ID == 0 {
		return loginVO, result.ERROR_USER_NOT_EXIST
	}

	// 检查密码是否正确
	if !utils.Encryptor.BcryptCheck(password, userAuth.Password) {
		return loginVO, result.ERROR_PASSWORD_WRONG
	}

	fmt.Println(userAuth)

	// 获取用户IP地址
	ipAddress := utils.IpUtil.GetIpAddress(c)
	browser, os := "unknown", "unknown"

	if userAgent := utils.IpUtil.GetUserAgent(c); userAgent != nil {
		browser = userAgent.Name + " " + userAgent.Version.String()
		os = userAgent.OS + " " + userAgent.OSVersion.String()
	}

	uuid := utils.Encryptor.MD5(ipAddress + browser + os)
	token, err := utils.GetJWT().GenToken(userAuth.ID, "1", uuid)

	if err != nil {
		utils.Logger.Info("登录时生成 Token 错误: ", zap.Error(err))
		return loginVO, result.ERROR_TOKEN_CREATE
	}

	loginVO = utils.CopyProperties[vo.LoginVO](userAuth)

	loginVO.Token = token

	// 保存用户信息到 Session 和 Redis 中
	// session := sessions.Default(c)
	sessionInfoStr := utils.Json.Marshal(dto.SessionInfo{loginVO})
	// session.Set(KEY_USER+uuid, sessionInfoStr)
	utils.Redis.Set(KEY_USER+uuid, sessionInfoStr, time.Duration(config.Cfg.Session.MaxAge)*time.Second)

	return loginVO, code
}

// 注册用户
func (s *UserService) Register(userDto req.Register) (code int) {
	// 检查验证码是否正确
	//if userDto.Code != utils.Redis.GetVal("code:"+userDto.Username) {
	//	return result.ERROR_VERIFICATION_CODE
	//}

	if exist := checkUserExistByUsername(userDto.Username); exist {
		return result.ERROR_USER_NAME_USED
	}

	dao.Create(&model.UserAuth{
		Username:      userDto.Username,
		Password:      utils.Encryptor.BcryptHash(userDto.Password),
		LoginType:     1,
		LastLoginTime: time.Now(),
	})

	return result.OK
}

func (s *UserService) GetUserInfo(id int) vo.UserInfoVO {
	var userInfo model.UserAuth
	dao.GetOne(&userInfo, "id", id)
	data := utils.CopyProperties[vo.UserInfoVO](userInfo)
	return data
}

// 检查用户名是否已存在
func checkUserExistByUsername(username string) bool {
	user := dao.GetOne(model.UserAuth{}, "username = ?", username)
	return user.ID != 0
}
