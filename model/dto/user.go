package dto

import "wow-admin/model/vo"

// Session 信息: 记录用户详细信息 + 是否被强退
type SessionInfo struct {
	vo.LoginVO
}

// 用户详细信息: 仅用于在后端内部进行传输
type UserDetailDTO struct {
	vo.LoginVO
}
