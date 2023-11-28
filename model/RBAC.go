package model

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
