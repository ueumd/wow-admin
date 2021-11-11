package dao

import (
	"bytes"
	"errors"
	"github.com/jmoiron/sqlx"
	"strings"
	"wow-admin/global"
	"wow-admin/model"
)
/**
用户操作
 */
var LogOperationInfosDao = logOperationInfosDao{}

type logOperationInfosDao struct{}

// 根据【自增ID】查询【操作流水表】表中是否存在相关记录
func (d *logOperationInfosDao) Exist(id int) (bool, error) {
	rows, err := global.DB.Queryx("select count(0) Count from log_operation_infos where id=?", id)
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

// 插入单条记录到【操作流水表】表中
func (d *logOperationInfosDao) Insert(m *model.LogOperationInfosModel) (int64, error) {
	insertStr := d.GetInsertItemString()
	result, err := global.DB.Exec("insert into log_operation_infos( "+insertStr+") values(?,?,?,?,?)", m.UserId, m.Phone, m.OpType, m.OpContent, m.IsDel)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【自增ID】修改【操作流水表】表的单条记录
func (d *logOperationInfosDao) Update(m *model.LogOperationInfosModel) (bool, error) {
	updateStr := d.GetUpdateItemString()
	result, err := global.DB.Exec("update log_operation_infos set "+updateStr+" where id=?", m.UserId, m.Phone, m.OpType, m.OpContent, m.IsDel, m.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【自增ID】软删除【操作流水表】表中的单条记录
func (d *logOperationInfosDao) Delete(id int) (bool, error) {
	result, err := global.DB.Exec("update log_operation_infos set isDel=1 where id=?", id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【自增ID】数组软删除【操作流水表】表中的多条记录
func (d *logOperationInfosDao) DeleteIn(ids []int) (count int64, err error) {
	if len(ids) <= 0 {
		return count, errors.New("ids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update log_operation_infos set isDel=1")
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

// 根据【自增ID】查询【操作流水表】表中的单条记录
func (d *logOperationInfosDao) Get(id int) (log_operation_infos model.LogOperationInfosModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select  "+selectStr+"  from log_operation_infos where id=?", id)
	if err != nil {
		return log_operation_infos, err
	}
	defer rows.Close()
	log_operation_infoss, err := d._RowsToArray(rows)
	if err != nil {
		return log_operation_infos, err
	}
	if len(log_operation_infoss) <= 0 {
		return log_operation_infos, err
	}
	return log_operation_infoss[0], nil
}

// 根据【自增ID】数组查询【操作流水表】表中的多条记录
func (d *logOperationInfosDao) GetIn(ids []int) (log_operation_infoss []model.LogOperationInfosModel, err error) {
	if len(ids) <= 0 {
		return log_operation_infoss, errors.New("ids is empty")
	}
	sql_str := bytes.Buffer{}
	selectStr := d.GetSelectItemString()
	sql_str.WriteString("select " + selectStr + "  from ")
	sql_str.WriteString("log_operation_infos")
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
		return log_operation_infoss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据【手机号】查询【操作流水表】表中的多条记录，使用索引【idx_phone,】
func (d *logOperationInfosDao) GetByPhone(phone string) (log_operation_infoss []model.LogOperationInfosModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from log_operation_infos force index(idx_phone) where phone=? and isDel = 0", phone)
	if err != nil {
		return log_operation_infoss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据【用户id】查询【操作流水表】表中的多条记录，使用索引【idx_userId,】
func (d *logOperationInfosDao) GetByUserId(userId int) (log_operation_infoss []model.LogOperationInfosModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from log_operation_infos force index(idx_userId) where userId=? and isDel = 0", userId)
	if err != nil {
		return log_operation_infoss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【操作流水表】表总记录数
func (d *logOperationInfosDao) GetRowCount() (count int, err error) {
	rows, err := global.DB.Queryx("select count(0) Count from log_operation_infos where isDel = 0")
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

// 根据【手机号】查询【操作流水表】表总记录数，使用索引【idx_phone,】
func (d *logOperationInfosDao) GetRowCountByPhone(phone string) (count int, err error) {
	rows, err := global.DB.Queryx("select count(0) Count from log_operation_infos force index(idx_phone) where phone=? and isDel = 0", phone)
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

// 根据【用户id】查询【操作流水表】表总记录数，使用索引【idx_userId,】
func (d *logOperationInfosDao) GetRowCountByUserId(userId int) (count int, err error) {
	rows, err := global.DB.Queryx("select count(0) Count from log_operation_infos force index(idx_userId) where userId=? and isDel = 0", userId)
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

// 分页查询【操作流水表】表的记录
func (d *logOperationInfosDao) GetRowList(PageIndex, PageSize int) (log_operation_infoss []model.LogOperationInfosModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from log_operation_infos where isDel = 0 limit ?,?", (PageIndex-1)*PageSize, PageSize)
	if err != nil {
		return log_operation_infoss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据【手机号】分页查询【操作流水表】表的记录，使用索引【idx_phone,】
func (d *logOperationInfosDao) GetRowListByPhone(phone string, PageIndex, PageSize int) (log_operation_infoss []model.LogOperationInfosModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from log_operation_infos force index(idx_phone) where phone=? and isDel = 0 limit ?,?", phone, (PageIndex-1)*PageSize, PageSize)
	if err != nil {
		return log_operation_infoss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据【用户id】分页查询【操作流水表】表的记录，使用索引【idx_userId,】
func (d *logOperationInfosDao) GetRowListByUserId(userId int, PageIndex, PageSize int) (log_operation_infoss []model.LogOperationInfosModel, err error) {
	selectStr := d.GetSelectItemString()
	rows, err := global.DB.Queryx("select "+selectStr+"  from log_operation_infos force index(idx_userId) where userId=? and isDel = 0 limit ?,?", userId, (PageIndex-1)*PageSize, PageSize)
	if err != nil {
		return log_operation_infoss, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 解析【操作流水表】表记录
func (d *logOperationInfosDao) _RowsToArray(rows *sqlx.Rows) (models []model.LogOperationInfosModel, err error) {
	for rows.Next() {
		mo := model.LogOperationInfosModel{}
		err = rows.StructScan(&mo)
		if err != nil {
			return models, err
		}
		models = append(models, mo)
	}
	return models, err
}

// 解析【操作流水表】表记录
// id, userId, phone, opType, opContent, isDel, createTime, updateTime
func (d *logOperationInfosDao) GetSelectItemString() string {
	selectStr := "id, userId, phone, opType, opContent, isDel, createTime, updateTime"
	return selectStr
}

// 解析【操作流水表】表记录
func (d *logOperationInfosDao) GetUpdateItemString() string {
	updateStr := "userId=?, phone=?, opType=?, opContent=?, isDel=?"
	return updateStr
}

// 解析【操作流水表】表记录
func (d *logOperationInfosDao) GetInsertItemString() string {
	return "userId,phone,opType,opContent,isDel"
}
