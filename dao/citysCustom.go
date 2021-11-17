package dao

import (
	"wow-admin/model"
	"wow-admin/global"
)

// 根据levelType查找
func (d *citysDao) GetByCityLevelType(levelType string) (cityss []model.CitysModel, err error) {
	rows, err := global.DB.Queryx("select id, cityName, parentId, shortName, levelType, cityCode, zipCode, mergerName, longitude, latitude, pinyin, isDel, createTime, updateId, updateTime from citys where levelType=? and isDel = 0 order by id", levelType)
	if err != nil {
		return cityss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据levelType查找
func (d *citysDao) GetByCityLevelTypeAndParentId(levelType string, parentId int) (cityss []model.CitysModel, err error) {
	rows, err := global.DB.Queryx("select id, cityName, parentId, shortName, levelType, cityCode, zipCode, mergerName, longitude, latitude, pinyin, isDel, createTime, updateId, updateTime from citys where levelType=? and parentId=? and isDel = 0", levelType, parentId)
	if err != nil {
		return cityss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}
