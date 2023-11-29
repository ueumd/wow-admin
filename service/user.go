package service

import (
	"github.com/gin-gonic/gin"
	"time"
	"wow-admin/dao"
	"wow-admin/model"
	"wow-admin/model/dto"
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

	return loginVO, code
}

// 注册用户
func (s *UserService) Register(userDto dto.Register) (code int) {

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

func checkUserExistByUsername(username string) bool {
	user := dao.GetOne(model.UserAuth{}, "username = ?", username)
	return user.ID != 0
}
