package dao

import (
	"github.com/jmoiron/sqlx"
	"wow-admin/global"
	"wow-admin/model"
)

var WechatUserDao = wechatUserDao{}

type wechatUserDao struct{}



// 根据【phone】查询【用户表】表中的多条记录，使用索引【idx_phone,】
func (d *wechatUserDao) GetByPhone(phone string) (wechat_users []model.WechatUserModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from wechat_user force index(idx_phone) where phone=? and isDel = 0", phone)
	if err != nil {
		return wechat_users, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}


// 根据【id】查询【微信用户表】表中的单条记录
func (d *wechatUserDao) Get(id int) (wechat_user model.WechatUserModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select  "+selectStr+"  from wechat_user where id=?", id)
	if err != nil {
		return wechat_user, err
	}
	defer rows.Close()
	wechat_users, err := d._RowsToArray(rows)
	if err != nil {
		return wechat_user, err
	}
	if len(wechat_users) <= 0 {
		return wechat_user, err
	}
	return wechat_users[0], nil
}

// 解析【用户表】表记录
func (d *wechatUserDao) _RowsToArray(rows *sqlx.Rows) (models []model.WechatUserModel, err error) {
	for rows.Next() {
		mo := model.WechatUserModel{}
		err = rows.StructScan(&mo)
		if err != nil {
			return models, err
		}
		models = append(models, mo)
	}
	return models, err
}

// 解析【用户表】表记录
// id, unionid, openid, phone, nickname, gender, headImgUrl, password, email, sketch, location, longitude, latitude, isDel, createTime, updateId, updateTime
func (d *wechatUserDao) GetSelectItemString() string {
	selectStr := "id, unionid, openid, phone, nickname, gender, headImgUrl, password, email, sketch, location, longitude, latitude, isDel, createTime, updateId, updateTime"
	return selectStr
}

// 解析【用户表】表记录
func (d *wechatUserDao) GetUpdateItemString() string {
	updateStr := "unionid=?, openid=?, phone=?, nickname=?, gender=?, headImgUrl=?, password=?, email=?, sketch=?, location=?, longitude=?, latitude=?, isDel=?, updateId=?"
	return updateStr
}

// 解析【用户表】表记录
func (d *wechatUserDao) GetInsertItemString() string {
	return "unionid,openid,phone,nickname,gender,headImgUrl,password,email,sketch,location,longitude,latitude,isDel,updateId"
}
