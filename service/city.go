package service

import (
	"fmt"
	"sort"
	"strings"
	"wow-admin/model/vo"
	"wow-admin/utils/cast"
)

type CityService struct {
}

func (*CityService) GetAllSupportCityList() (*vo.AllSupportCityVO, error) {
	allSupportCityVO := &vo.AllSupportCityVO{}

	cityListVO := make([]vo.CityVO, 0)

	cityListModel := cityDao.GetCityByLevelType("2")

	for _, c := range cityListModel {
		cityName := c.ShortName
		if strings.TrimSpace(cityName) == "" {
			cityName = c.CityName
		}

		if strings.TrimSpace(cityName) == " " {
			continue
		}

		cityListVO = append(cityListVO, vo.CityVO{
			Name:      c.ShortName,
			Id:        fmt.Sprintf("%d", c.ID),
			LevelType: c.LevelType,
		})
	}

	sort.Slice(cityListVO, func(i, j int) bool {
		pinyinI := cast.GetCityPinYin(cityListVO[i].Name)
		pinyinJ := cast.GetCityPinYin(cityListVO[j].Name)

		if pinyinI > pinyinJ {
			return false
		}

		return true
	})

	cityList := make([]*vo.LetterCityVO, 0)
	cityMaps := make(map[string]*vo.LetterCityVO)
	for _, cc := range cityListVO {
		letter := cast.GetCityFirstLetter(cc.Name)
		v, ok := cityMaps[letter]
		if ok {
			v.Data = append(v.Data, cc)
		} else {
			cTemp := &vo.LetterCityVO{}
			cTemp.Letter = letter
			cTemp.Data = make([]vo.CityVO, 0)
			cTemp.Data = append(cTemp.Data, cc)
			cityList = append(cityList, cTemp)
			cityMaps[letter] = cTemp
		}
	}
	allSupportCityVO.CityList = cityList

	hotCityList := make([]vo.CityVO, 0)
	hotCityList = append(hotCityList, vo.CityVO{
		Name:      "北京",
		Id:        "110100",
		LevelType: "2",
	})

	hotCityList = append(hotCityList, vo.CityVO{
		Name:      "上海",
		Id:        "430100",
		LevelType: "2",
	})

	hotCityList = append(hotCityList, vo.CityVO{
		Name:      "杭州",
		Id:        "330100",
		LevelType: "2",
	})

	hotCityList = append(hotCityList, vo.CityVO{
		Name:      "长沙",
		Id:        "430100",
		LevelType: "2",
	})

	hotCityList = append(hotCityList, vo.CityVO{
		Name:      "苏州",
		Id:        "320500",
		LevelType: "2",
	})

	hotCityList = append(hotCityList, vo.CityVO{
		Name:      "广州",
		Id:        "440100",
		LevelType: "2",
	})

	hotCityList = append(hotCityList, vo.CityVO{
		Name:      "深圳",
		Id:        "440300",
		LevelType: "2",
	})

	hotCityList = append(hotCityList, vo.CityVO{
		Name:      "南京",
		Id:        "320100",
		LevelType: "2",
	})

	hotCityList = append(hotCityList, vo.CityVO{
		Name:      "天津",
		Id:        "120100",
		LevelType: "2",
	})

	allSupportCityVO.HotCityList = hotCityList
	return allSupportCityVO, nil
}
