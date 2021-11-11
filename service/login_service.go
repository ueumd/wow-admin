package service

import (
	"fmt"
	"github.com/ueumd/logger"
	"time"
	"wow-admin/dao"
	"wow-admin/global"
	"wow-admin/global/cerrors"
	"wow-admin/model"
	"wow-admin/model/loginreq"
)


type LoginService struct {
	checkCodeService CommService
	logService       LogService
}

// 登录
func (l *LoginService) PhoneLogin(phoneLoginReq *loginreq.PhoneLoginRequest) (map[string]interface{}, error) {
	var err error

	//密码, 验证码 先省略

	//if phoneLoginReq.Password == "" {
	//	return nil,cerrors.New400ServiceError("请输入密码")
	//}

	//if phoneLoginReq.CheckCode == "" {
	//	return nil,cerrors.New400ServiceError("请填写验证码")
	//}

	err = l.checkCodeService.CheckPhoneCheckCode(CodeTypeLogin, phoneLoginReq.Phone, phoneLoginReq.CheckCode) //checkCode
	if err != nil {
		return nil,err
	}
	var wechatUser model.WechatUserModel
	wechatUserList, err := dao.WechatUserDao.GetByPhone(phoneLoginReq.Phone)
	if err != nil {
		logger.ErrorF("SQLERROR phoneLogin select user error %v", err)
		return nil, cerrors.New500ServiceError()
	}

	if len(wechatUserList) < 1 { //用户不存在
		logger.Errorln("用户或密码不正确")
		return nil, cerrors.NewServiceErrorCode(cerrors.ErrLoginErr, "")
	}else {
		wechatUser = wechatUserList[0]
	}

	// 记录用户登录日志信息
	l.logService.InsertLoginLog(wechatUser.Id, fmt.Sprintf("%s", wechatUser.Phone), map[string]interface{}{
		"result":    "SUCCESS",
		"loginType": "Phone",
		"param":     phoneLoginReq,
	})


	redisKey := fmt.Sprintf("login:userId:%d", wechatUser.Id)

	// 更新token时间
	if tk, err := global.RedisClient.Get(redisKey).Result(); err == nil && tk != "" {
		global.RedisClient.Set(redisKey, tk, 1 * 2000)
		return map[string]interface{}{
			"token": tk,
		}, nil
	}

	// 设置新token
	_, token := GenerateToken(wechatUser.Id, wechatUser.Unionid, fmt.Sprintf("%d", wechatUser.Phone), "wechat")
	global.RedisClient.Set(redisKey, token, 7*24*time.Hour)
	return  map[string]interface{}{
		"token": token,
	}, nil
}

// 获取token
func (l *LoginService) GetLoginToken(userId int) (string,error) {
	redisKey := fmt.Sprintf("login:userId:%d", userId)
	var tk string
	var err error
	if tk, err = global.RedisClient.Get(redisKey).Result(); err == nil && tk != "" {
		global.RedisClient.Set(redisKey, tk, 7*24*time.Hour)
		return tk, nil
	}

	return "", err
}


func (l *LoginService) GetOrGenerateUserToken(user model.WechatUserModel) string {
	redisKey := fmt.Sprintf("login:userId:%d", user.Id)
	if tk, err := global.RedisClient.Get(redisKey).Result(); err == nil && tk != "" {
		global.RedisClient.Set(redisKey, tk, 7*24*time.Hour)
		return tk
	}

	_, token := GenerateToken(user.Id, user.Openid, fmt.Sprintf("%s", user.Phone), "wechat")
	global.RedisClient.Set(redisKey, token, 7*24*time.Hour)
	return token
}






