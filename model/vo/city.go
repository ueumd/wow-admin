package vo

type CityVO struct {
	Name      string    `json:"name"`
	Id        string    `json:"id"`
	LevelType string    `json:"levelType"`
	Children  []*CityVO `json:"children"`
}

type LetterCityVO struct {
	Letter string   `json:"letter"`
	Data   []CityVO `json:"data"`
}

type AllSupportCityVO struct {
	CityList    []*LetterCityVO `json:"cityList"`
	HotCityList []CityVO        `json:"hotCityList"`
}
