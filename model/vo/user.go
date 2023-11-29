package vo

import "time"

// 登录VO
type LoginVO struct {
	ID            int       `json:"id"`
	UserInfoId    int       `json:"userInfoId"`
	Username      string    `json:"username"`
	LastLoginTime time.Time `json:"last_login_time"`
	LoginType     int       `json:"login_type"`
	Token         string    `json:"token"`
}
