package userrep

type RegisterInfoResponse struct {
	Name      string                 `json:"name"`
	Id        string                 `json:"id"`
	LevelType string                 `json:"levelType"`
	Longitude string                 `json:"longitude"`
	Latitude  string                 `json:"latitude"`
	Children  []RegisterInfoResponse `json:"children"`
}
