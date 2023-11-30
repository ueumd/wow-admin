package vo

import "time"

// 登录VO
type LoginVO struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	LastLoginTime time.Time `json:"lastLoginTime"`
	IpAddress     string    `json:"ipAddress"`
	IpSource      string    `json:"ipSource"`
	LoginType     int       `json:"loginType"`
	Token         string    `json:"token"`
}

// 用户信息 VO
type UserInfoVO struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	LoginType int    `json:"loginType"`
}
