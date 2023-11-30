package model

import "time"

/**
Role-Based Access Control
*/

// 菜单
type Menu struct {
	Universal `mapstructure:",squash"`
	Name      string `gorm:"type:varchar(20);comment:菜单名" json:"name"`
	Path      string `gorm:"type:varchar(50);comment:菜单路径" json:"path"`
	Component string `gorm:"type:varchar(50);comment:组件" json:"component"`
	Icon      string `gorm:"type:varchar(50);comment:菜单图标" json:"icon"`
	ParentId  int    `gorm:"comment:父菜单id" json:"parent_id"`
	OrderNum  int8   `gorm:"type:tinyint;default:0;comment:显示排序" json:"order_num"`
	IsHidden  int8   `gorm:"type:tinyint(1);default:0;comment:是否隐藏(0-否 1-是)" json:"is_hidden"`
	KeepAlive int8   `gorm:"type:tinyint(1);default:1" json:"keep_alive"`
	Redirect  string `gorm:"type:varchar(50);comment:跳转路径" json:"redirect"`
}

// 用户账户信息
type UserAuth struct {
	Universal
	Username      string    `gorm:"type:varchar(50);comment:用户名" json:"username"`
	Password      string    `gorm:"type:varchar(100)" json:"password"`
	LoginType     int       `gorm:"type:tinyint(1);comment:登录类型" json:"loginType"`
	IpAddress     string    `gorm:"type:varchar(20);comment:登录IP地址" json:"ipAddress"`
	IpSource      string    `gorm:"type:varchar(50);comment:IP来源" json:"ip_source"`
	LastLoginTime time.Time `gorm:"comment:上次登录时间" json:"LastLoginTime"`
}
