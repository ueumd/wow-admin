package loginreq

type PhoneLoginRequest struct {
	Phone string `json:"phone"`
	CheckCode string `json:"checkCode"` //验证码
	Password  string `json:"password"`  //密码
}
