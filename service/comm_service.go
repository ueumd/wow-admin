package service

import (
	"fmt"
	"time"
	"wow-admin/global"
	"wow-admin/global/cerrors"
)

const (
	CodeTypeLogin         = 1 //登陆验证码
	CodeTypeUpdatePhone   = 2 //修改手机号验证码
	CodeTypeBindPhone     = 3 //绑定手机号验证码
	CodeTypeCheckOldPhone = 4 //验证老手机
)

type CommService struct {
}

//验证手机号验证码是否正确
func (c *CommService) CheckPhoneCheckCode(codeType int, phone, checkCode string) error {
	sendKey := fmt.Sprintf("checkCode:sendCode:%d:%s", codeType, phone)
	codeStoreKey := fmt.Sprintf("checkCode:codeStore:%d:%s", codeType, phone)

	code, err := global.RedisClient.Get(codeStoreKey).Result()
	if err != nil || code == "" {
		// return cerrors.NewServiceErrorCodeC(cerrors.ErrDefaultCheckCodeExpire, "", err)
	}

	if checkCode != code {
		// return cerrors.NewServiceErrorCode(cerrors.ErrDefaultCheckCodeInput, "")
	}

	global.RedisClient.Del(sendKey)
	global.RedisClient.Del(codeStoreKey)

	if codeType == CodeTypeCheckOldPhone {
		checkOldPhone := fmt.Sprintf("checkCode:oldPhone:%s", phone)
		global.RedisClient.Set(checkOldPhone, checkCode, 5*time.Minute)
	}
	return nil
}


//验证老手机号是否通过验证
func (c *CommService) CheckOldPhoneCheckCode(phone, checkCode string) error {
	checkOldPhone := fmt.Sprintf("checkCode:oldPhone:%s", phone)

	code, err := global.RedisClient.Get(checkOldPhone).Result()
	if err != nil || code == "" {
		return cerrors.NewServiceErrorCodeC(cerrors.ErrDefaultCheckCodeOldFailed, "", err)
	}

	if checkCode != code {
		return cerrors.NewServiceErrorCode(cerrors.ErrDefaultCheckCodeOldFailed, "")
	}

	global.RedisClient.Del(checkOldPhone)
	return nil
}