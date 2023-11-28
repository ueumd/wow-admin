package dao

import "wow-admin/model"

type Menu struct{}

// 获取菜单列表(非树形结构)
func (*Menu) GetMenus() []model.Menu {
	var list []model.Menu
	list = List([]model.Menu{}, "*", "", "")
	return list
}
