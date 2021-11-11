package model

import (
	"time"
)

// 微信用户表
type WechatUserModel struct {
	Id         int       `db:"id"`
	Unionid    string    `db:"unionid"`
	Openid     string    `db:"openid"`
	Phone      string    `db:"phone"`
	Nickname   string    `db:"nickname"`   //呢称
	Gender     int8      `db:"gender"`     //1男，2:女, 3:保密
	HeadImgUrl string    `db:"headImgUrl"` //头像地址
	Password   string    `db:"password"`   //密码
	Email      string    `db:"email"`      //邮箱
	Sketch     string    `db:"sketch"`     //简述
	Location   string    `db:"location"`   //当前位置
	Longitude  string    `db:"longitude"`  //经度
	Latitude   string    `db:"latitude"`   //纬度
	IsDel      int8      `db:"isDel"`      //1已删除 0未删除
	CreateTime time.Time `db:"createTime"`
	UpdateId   int       `db:"updateId"`
	UpdateTime time.Time `db:"updateTime"`
}
