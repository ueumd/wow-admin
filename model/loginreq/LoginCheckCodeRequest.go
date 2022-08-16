package loginreq

//获取登陆验证码
type LoginCheckCodeRequest struct {
	Phone string `json:"phone"` //
	ImgCheckCode string `json:"imageCode"` //图形验证码
	DeviceId string `json:"deviceId"`
}