package userrep

type AreaResponse struct {
	Name      string          `json:"name"`
	Id        string          `json:"id"`
	LevelType string          `json:"levelType"`
	Children  []*AreaResponse `json:"children"`
}

type LetterAreaResp struct {
	Letter string         `json:"letter"`
	Data   []AreaResponse `json:"data"`
}

type AllSupportAreaResponse struct {
	CityList []*LetterAreaResp `json:"cityList"`
	HotCitys []AreaResponse    `json:"hotCityList"`
}
