package service

import (
	"wow-admin/model"
	"wow-admin/model/vo"
)

type MenuService struct{}

// 获取菜单列表(树形)
func (s *MenuService) GetTreeList() []vo.MenuVo {

	resMenuVos := make([]vo.MenuVo, 0)

	menuList := menuDao.GetMenus()

	for _, item := range menuList {
		menuVo := s.menu2MenuVo(item)
		resMenuVos = append(resMenuVos, menuVo)
	}

	return resMenuVos
}

// model.Menu => vo.MenuVo
func (*MenuService) menu2MenuVo(menu model.Menu) vo.MenuVo {
	return vo.MenuVo{
		ID:        menu.ID,
		Name:      menu.Name,
		Path:      menu.Path,
		Component: menu.Component,
		Icon:      menu.Icon,
		CreatedAt: menu.CreatedAt,
		OrderNum:  menu.OrderNum,
		IsHidden:  int(menu.IsHidden),
		ParentId:  menu.ParentId,
		Redirect:  menu.Redirect,
		KeepAlive: menu.KeepAlive,
		Children:  make([]vo.MenuVo, 0),
	}
}
