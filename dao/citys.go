package dao

import (
	"bytes"
	"errors"
	"github.com/jmoiron/sqlx"
	"wow-admin/global"
	"wow-admin/model"
	"strings"
)

var CitysDao = citysDao{}

type citysDao struct{}

// 根据【城市ID】查询【城市信息表】表中是否存在相关记录
func (d *citysDao) Exist(id int) (bool, error) {
	rows, err := global.DB.Queryx("select count(0) Count from citys where id=?", id)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	count := 0
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return false, err
		}
		return count > 0, nil
	}
	return false, nil
}

// 插入单条记录到【城市信息表】表中
func (d *citysDao) Insert(m *model.CitysModel) (bool, error) {
	insertStr := d.GetInsertItemString()
	result, err := global.DB.Exec("insert into citys( "+insertStr+") values(?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Id, m.CityName, m.ParentId, m.ShortName, m.LevelType, m.CityCode, m.ZipCode, m.MergerName, m.Longitude, m.Latitude, m.Pinyin, m.IsDel, m.UpdateId)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【城市ID】修改【城市信息表】表的单条记录
func (d *citysDao) Update(m *model.CitysModel) (bool, error) {
	updateStr := d.GetUpdateItemString()
	result, err := global.DB.Exec("update citys set "+updateStr+" where id=?", m.CityName, m.ParentId, m.ShortName, m.LevelType, m.CityCode, m.ZipCode, m.MergerName, m.Longitude, m.Latitude, m.Pinyin, m.IsDel, m.UpdateId, m.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 插入或修改【城市信息表】表的单条记录
func (d *citysDao) InsertUpdate(m *model.CitysModel) (bool, error) {
	insertStr := d.GetInsertItemString()
	result, err := global.DB.Exec("insert into citys("+insertStr+") values(?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE cityName=?,parentId=?,shortName=?,levelType=?,cityCode=?,zipCode=?,mergerName=?,longitude=?,latitude=?,pinyin=?,isDel=?,updateId=?", m.Id, m.CityName, m.ParentId, m.ShortName, m.LevelType, m.CityCode, m.ZipCode, m.MergerName, m.Longitude, m.Latitude, m.Pinyin, m.IsDel, m.UpdateId, m.CityName, m.ParentId, m.ShortName, m.LevelType, m.CityCode, m.ZipCode, m.MergerName, m.Longitude, m.Latitude, m.Pinyin, m.IsDel, m.UpdateId)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【城市ID】软删除【城市信息表】表中的单条记录
func (d *citysDao) Delete(id int) (bool, error) {
	result, err := global.DB.Exec("update citys set isDel=1 where id=?", id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【城市ID】数组软删除【城市信息表】表中的多条记录
func (d *citysDao) DeleteIn(ids []int) (count int64, err error) {
	if len(ids) <= 0 {
		return count, errors.New("ids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update citys set isDel=1")
	sql_str.WriteString(" where id in(")
	question_mark := strings.Repeat("?,", len(ids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(ids))
	for _, v := range ids {
		vals = append(vals, v)
	}
	result, err := global.DB.Exec(sql_str.String(), vals...)
	if err != nil {
		return count, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return count, err
	}
	return affected, nil
}

// 根据【城市ID】查询【城市信息表】表中的单条记录
func (d *citysDao) Get(id int) (citys model.CitysModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select  "+selectStr+"  from citys where id=?", id)
	if err != nil {
		return citys, err
	}
	defer rows.Close()
	cityss, err := d._RowsToArray(rows)
	if err != nil {
		return citys, err
	}
	if len(cityss) <= 0 {
		return citys, err
	}
	return cityss[0], nil
}

// 根据【城市ID】数组查询【城市信息表】表中的多条记录
func (d *citysDao) GetIn(ids []int) (cityss []model.CitysModel, err error) {
	if len(ids) <= 0 {
		return cityss, errors.New("ids is empty")
	}
	sql_str := bytes.Buffer{}
	selectStr := d.GetSelectItemString()
	sql_str.WriteString("select " + selectStr + "  from ")
	sql_str.WriteString("citys")
	sql_str.WriteString(" where id in(")
	param_keys := strings.Repeat("?,", len(ids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(ids))
	for _, v := range ids {
		vals = append(vals, v)
	}
	rows, err = global.DB.Queryx(sql_str.String(), vals...)
	if err != nil {
		return cityss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据【城市名称】查询【城市信息表】表中的多条记录，使用索引【idx_cityName,】
func (d *citysDao) GetByCityName(cityName string) (cityss []model.CitysModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from citys force index(idx_cityName) where cityName=? and isDel = 0", cityName)
	if err != nil {
		return cityss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据【父城市ID】查询【城市信息表】表中的多条记录，使用索引【idx_parentId,】
func (d *citysDao) GetByParentId(parentId int) (cityss []model.CitysModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from citys force index(idx_parentId) where parentId=? and isDel = 0", parentId)
	if err != nil {
		return cityss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【城市信息表】表总记录数
func (d *citysDao) GetRowCount() (count int, err error) {
	rows, err := global.DB.Queryx("select count(0) Count from citys where isDel = 0")
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

// 根据【城市名称】查询【城市信息表】表总记录数，使用索引【idx_cityName,】
func (d *citysDao) GetRowCountByCityName(cityName string) (count int, err error) {
	rows, err := global.DB.Queryx("select count(0) Count from citys force index(idx_cityName) where cityName=? and isDel = 0", cityName)
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

// 根据【父城市ID】查询【城市信息表】表总记录数，使用索引【idx_parentId,】
func (d *citysDao) GetRowCountByParentId(parentId int) (count int, err error) {
	rows, err := global.DB.Queryx("select count(0) Count from citys force index(idx_parentId) where parentId=? and isDel = 0", parentId)
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

// 分页查询【城市信息表】表的记录
func (d *citysDao) GetRowList(PageIndex, PageSize int) (cityss []model.CitysModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from citys where isDel = 0 limit ?,?", (PageIndex-1)*PageSize, PageSize)
	if err != nil {
		return cityss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据【父城市ID】分页查询【城市信息表】表的记录，使用索引【idx_parentId,】
func (d *citysDao) GetRowListByParentId(parentId int, PageIndex, PageSize int) (cityss []model.CitysModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from citys force index(idx_parentId) where parentId=? and isDel = 0 limit ?,?", parentId, (PageIndex-1)*PageSize, PageSize)
	if err != nil {
		return cityss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据【城市名称】分页查询【城市信息表】表的记录，使用索引【idx_cityName,】
func (d *citysDao) GetRowListByCityName(cityName string, PageIndex, PageSize int) (cityss []model.CitysModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from citys force index(idx_cityName) where cityName=? and isDel = 0 limit ?,?", cityName, (PageIndex-1)*PageSize, PageSize)
	if err != nil {
		return cityss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 解析【城市信息表】表记录
func (d *citysDao) _RowsToArray(rows *sqlx.Rows) (models []model.CitysModel, err error) {
	for rows.Next() {
		mo := model.CitysModel{}
		err = rows.StructScan(&mo)
		if err != nil {
			return models, err
		}
		models = append(models, mo)
	}
	return models, err
}

// 解析【城市信息表】表记录
// id, cityName, parentId, shortName, levelType, cityCode, zipCode, mergerName, longitude, latitude, pinyin, isDel, createTime, updateId, updateTime
func (d *citysDao) GetSelectItemString() string {
	selectStr := "id, cityName, parentId, shortName, levelType, cityCode, zipCode, mergerName, longitude, latitude, pinyin, isDel, createTime, updateId, updateTime"
	return selectStr
}

// 解析【城市信息表】表记录
func (d *citysDao) GetUpdateItemString() string {
	updateStr := "cityName=?, parentId=?, shortName=?, levelType=?, cityCode=?, zipCode=?, mergerName=?, longitude=?, latitude=?, pinyin=?, isDel=?, updateId=?"
	return updateStr
}

// 解析【城市信息表】表记录
func (d *citysDao) GetInsertItemString() string {
	return "id,cityName,parentId,shortName,levelType,cityCode,zipCode,mergerName,longitude,latitude,pinyin,isDel,updateId"
}
