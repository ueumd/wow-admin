package dao

import (
	"wow-admin/global"
)

var RegisterInfoDao = registerInfoDao{}

type registerInfoDao struct{}

// 根据【手机号码】查询【注册信息表】表总记录数，使用索引【idx_phone,】
func (d *registerInfoDao) GetRowCountByPhone(phone string) (count int, err error) {
	rows, err := global.DB.Queryx("select count(0) Count from register_info force index(idx_phone) where phone=? and isDel = 0", phone)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	return -1, nil
}