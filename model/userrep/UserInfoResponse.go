package userrep

type UserInfoResponse struct {
	UserId               int    `json:"userId"`
	HeadImgURL           string `json:"headImgUrl"`
	NickName             string `json:"nickName"`
	Phone                string `json:"phone"`
	CallName             string `json:"callName"`
	ChildAge             string `json:"childAge"`
	ChildName            string `json:"childName"`
	ChildSex             string `json:"childSex"`
	CityCode             string `json:"cityCode"`
	CityName             string `json:"cityName"`
	ParentWho            int    `json:"parentWho"`
	ParentWhoTxt         string `json:"parentWhoTxt"`
	Budget               string `json:"budget"`
	LearningLesson       string `json:"learningLesson"`
	StudyWishSkill       string `json:"studyWishSkill"`
	GivePriorityToLesson string `json:"givePriorityToLesson"`
	IntentionStudyType   string `json:"intentionStudyType"`
}
