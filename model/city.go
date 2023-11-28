package model

type City struct {
	Universal
	CityName   string `gorm:"type:varchar(30);not null" json:"cityName"`   //城市名称
	ParentId   int    `gorm:"type:int;not null" json:"parentId"`           //父城市ID
	ShortName  string `gorm:"type:varchar(30);not null" json:"shortName"`  //简介
	LevelType  string `gorm:"type:char" json:"levelType"`                  //城市等级
	CityCode   string `gorm:"type:varchar(8);not null" json:"cityCode"`    //城市代码
	ZipCode    string `gorm:"type:varchar(6);not null" json:"zipCode"`     //邮编
	MergerName string `gorm:"type:varchar(50);not null" json:"mergerName"` //城市名称
	Longitude  string `gorm:"type:varchar(15);not null" json:"longitude"`  //经度
	Latitude   string `gorm:"type:varchar(15);not null" json:"latitude"`   //纬度
	Pinyin     string `gorm:"type:varchar(30);not null" json:"pinyin"`     //简拼
	IsDel      int8   `gorm:"type:tinyint;not null" json:"isDel"`          //1 已删除  0 未删除
	UpdateId   int    `gorm:"type:int" json:"updateId"`                    //修改人Id
}
