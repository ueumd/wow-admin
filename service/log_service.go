package service

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/ueumd/logger"
	"wow-admin/dao"
	"wow-admin/model"
)

type OpType string

const (
	Login      OpType = "login"
	BindPhone  OpType = "bindPhone"
)

type LogService struct {}

func (l *LogService) InsertLoginLog(usrId int, phone string, linfo interface{}) {
	lgModel := &model.LogOperationInfosModel{}
	lgModel.UserId = usrId
	lgModel.Phone = phone
	lgModel.OpType = string(Login)
	bys,_ := jsoniter.Marshal(&linfo)
	lgModel.OpContent = string(bys)

	_,err := dao.LogOperationInfosDao.Insert(lgModel)
	if err != nil {
		logger.ErrorF("SQLERROR insert login log error:%v", err)
	}
}

func (l *LogService) InsertBindPhoneLog(phone string, linfo interface{}) {
	lgModel := &model.LogOperationInfosModel{}
	lgModel.Phone = phone
	lgModel.OpType = string(BindPhone)
	bys,_ := jsoniter.Marshal(&linfo)
	lgModel.OpContent = string(bys)

	dao.LogOperationInfosDao.Insert(lgModel)
}

