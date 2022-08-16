package service

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"wow-admin/dao"
	"wow-admin/global"
	"wow-admin/global/cerrors"
	"wow-admin/model/userrep"
	"wow-admin/utils/cast"
)

const (
	CodeTypeLogin         = 1 //登陆验证码
	CodeTypeUpdatePhone   = 2 //修改手机号验证码
	CodeTypeBindPhone     = 3 //绑定手机号验证码
	CodeTypeCheckOldPhone = 4 //验证老手机
)

type CommService struct {
}

//验证手机号验证码是否正确
func (c *CommService) CheckPhoneCheckCode(codeType int, phone, checkCode string) error {
	sendKey := fmt.Sprintf("checkCode:sendCode:%d:%s", codeType, phone)
	codeStoreKey := fmt.Sprintf("checkCode:codeStore:%d:%s", codeType, phone)

	code, err := global.RedisClient.Get(codeStoreKey).Result()
	if err != nil || code == "" {
		// return cerrors.NewServiceErrorCodeC(cerrors.ErrDefaultCheckCodeExpire, "", err)
	}

	if checkCode != code {
		// return cerrors.NewServiceErrorCode(cerrors.ErrDefaultCheckCodeInput, "")
	}

	global.RedisClient.Del(sendKey)
	global.RedisClient.Del(codeStoreKey)

	if codeType == CodeTypeCheckOldPhone {
		checkOldPhone := fmt.Sprintf("checkCode:oldPhone:%s", phone)
		global.RedisClient.Set(checkOldPhone, checkCode, 5*time.Minute)
	}
	return nil
}


//验证老手机号是否通过验证
func (c *CommService) CheckOldPhoneCheckCode(phone, checkCode string) error {
	checkOldPhone := fmt.Sprintf("checkCode:oldPhone:%s", phone)

	code, err := global.RedisClient.Get(checkOldPhone).Result()
	if err != nil || code == "" {
		return cerrors.NewServiceErrorCodeC(cerrors.ErrDefaultCheckCodeOldFailed, "", err)
	}

	if checkCode != code {
		return cerrors.NewServiceErrorCode(cerrors.ErrDefaultCheckCodeOldFailed, "")
	}

	global.RedisClient.Del(checkOldPhone)
	return nil
}

func (c *CommService) GetAllSupportCityList() (*userrep.AllSupportAreaResponse, error) {
	resp := &userrep.AllSupportAreaResponse{}
	allCitys := make([]userrep.AreaResponse, 0)

	citys2, err := dao.CitysDao.GetByCityLevelType("2")
	if err != nil {
		return resp, cerrors.New600ServiceErrorC(err)
	}

	for _, c := range citys2 {
		cityName := c.ShortName
		if strings.TrimSpace(cityName) == "" {
			cityName = c.CityName
		}

		if strings.TrimSpace(cityName) == " " {
			continue
		}

		allCitys = append(allCitys, userrep.AreaResponse{
			Name:      c.ShortName,
			Id:        fmt.Sprintf("%d", c.Id),
			LevelType: c.LevelType,
		})
	}

	sort.Slice(allCitys, func(i, j int) bool {
		pinyinI := cast.GetCityPinYin(allCitys[i].Name)
		pinyinJ := cast.GetCityPinYin(allCitys[j].Name)

		if pinyinI > pinyinJ {
			return false
		}

		return true
	})

	cityList := make([]*userrep.LetterAreaResp, 0)
	cityMaps := make(map[string]*userrep.LetterAreaResp)
	for _, cc := range allCitys {
		letter := cast.GetCityFirstLetter(cc.Name)
		v, ok := cityMaps[letter]
		if ok {
			v.Data = append(v.Data, cc)
		} else {
			cTemp := &userrep.LetterAreaResp{}
			cTemp.Letter = letter
			cTemp.Data = make([]userrep.AreaResponse, 0)
			cTemp.Data = append(cTemp.Data, cc)
			cityList = append(cityList, cTemp)
			cityMaps[letter] = cTemp
		}
	}
	resp.CityList = cityList

	hotCitys := make([]userrep.AreaResponse, 0)
	hotCitys = append(hotCitys, userrep.AreaResponse{
		Name:      "北京",
		Id:        "110100",
		LevelType: "2",
	})

	hotCitys = append(hotCitys, userrep.AreaResponse{
		Name:      "上海",
		Id:        "430100",
		LevelType: "2",
	})

	hotCitys = append(hotCitys, userrep.AreaResponse{
		Name:      "杭州",
		Id:        "330100",
		LevelType: "2",
	})

	hotCitys = append(hotCitys, userrep.AreaResponse{
		Name:      "长沙",
		Id:        "430100",
		LevelType: "2",
	})

	hotCitys = append(hotCitys, userrep.AreaResponse{
		Name:      "苏州",
		Id:        "320500",
		LevelType: "2",
	})

	hotCitys = append(hotCitys, userrep.AreaResponse{
		Name:      "广州",
		Id:        "440100",
		LevelType: "2",
	})

	hotCitys = append(hotCitys, userrep.AreaResponse{
		Name:      "深圳",
		Id:        "440300",
		LevelType: "2",
	})

	hotCitys = append(hotCitys, userrep.AreaResponse{
		Name:      "南京",
		Id:        "320100",
		LevelType: "2",
	})

	hotCitys = append(hotCitys, userrep.AreaResponse{
		Name:      "天津",
		Id:        "120100",
		LevelType: "2",
	})

	resp.HotCitys = hotCitys
	return resp, nil
}
