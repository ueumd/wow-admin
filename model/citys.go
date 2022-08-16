package model

import (
	"time"
)

// 城市信息表
type CitysModel struct {
	Id         int       `db:"id"`         //城市ID
	CityName   string    `db:"cityName"`   //城市名称
	ParentId   int       `db:"parentId"`   //父城市ID
	ShortName  string    `db:"shortName"`  //简介
	LevelType  string    `db:"levelType"`  //城市等级
	CityCode   string    `db:"cityCode"`   //城市代码
	ZipCode    string    `db:"zipCode"`    //邮编
	MergerName string    `db:"mergerName"` //城市名称
	Longitude  string    `db:"longitude"`  //经度
	Latitude   string    `db:"latitude"`   //纬度
	Pinyin     string    `db:"pinyin"`     //简拼
	IsDel      int8      `db:"isDel"`      //1 已删除  0 未删除
	CreateTime time.Time `db:"createTime"` //创建日期
	UpdateId   int       `db:"updateId"`   //修改人Id
	UpdateTime time.Time `db:"updateTime"` //修改日期
}
