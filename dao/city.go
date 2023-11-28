package dao

import "wow-admin/model"

type City struct {
}

// 根据levelType查找
func (d *City) GetCityByLevelType(levelType string) []model.City {
	var list []model.City
	list = List([]model.City{}, "*", "", "is_del = 0 and level_type=?", levelType)
	return list
}
