package service

import "wow-admin/dao"

// redis key

const (
	KEY_CODE = "code:" // 验证码
	KEY_USER = "user:" // 记录用户
	KEY_PAGE = "page"  // 页面封面
)

var (
	menuDao dao.Menu
	cityDao dao.City
)
