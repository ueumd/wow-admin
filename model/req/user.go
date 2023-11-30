package req

type Login struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

type UpdatePassword struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

type Register struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required,min=4,max=20" label:"密码"`
	Code     string `json:"code" validate:"required" label:"邮箱验证码"`
}
