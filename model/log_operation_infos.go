package model

import (
	"time"
)

// 操作流水表
type LogOperationInfosModel struct {
	Id         int       `db:"id"`     //自增ID
	UserId     int       `db:"userId"` //用户id
	Phone      string    `db:"phone"`  //手机号
	OpType     string    `db:"opType"` //操作类型   login   logout   updateInfo
	OpContent  string    `db:"opContent"`
	IsDel      int8      `db:"isDel"` //1已删除 0未删除
	CreateTime time.Time `db:"createTime"`
	UpdateTime time.Time `db:"updateTime"`
}
