package service

import (
	"wow-admin/dao"
	"wow-admin/global/cerrors"
	"wow-admin/model"
)

type UserBasicService struct {
}

func (u *UserBasicService) GetWeChatUser(userId int) (model.WechatUserModel, error) {
	user,err := dao.WechatUserDao.Get(userId)
	if err != nil {
		return user, cerrors.New600ServiceError()
	}

	if user.Id < 1 || user.IsDel == 1 {
		return user, cerrors.SResultError(cerrors.ErrUserNotExist, nil)
	}

	return user, nil
}