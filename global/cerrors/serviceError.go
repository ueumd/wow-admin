package cerrors

const (
	// 登录相关 10001000 开始
	ErrLoginCodeInput      = 10001001
	ErrLoginCodeRepeatSend = 10001002
	ErrLoginCodeSend       = 10001003
	ErrLoginCodeExpire     = 10001004
	ErrLoginNickName       = 10001005
	ErrLoginWxFailed       = 10001006
	//用户 10002000 开始
	ErrUserNotExist        = 100020001
	ErrUserWasDelete       = 100020002
	ErrUserUpdate          = 100020003
	ErrUserPhoneBinded     = 100020004
	ErrUserHasBindPhone    = 100020005
	ErrUserBindPhoneFailed = 100020006
	ErrUserHasNoPhone      = 100020007
	ErrUserPhoneIsNotNull  = 100020010

	//comm
	ErrCommLocationFailed = 100090003
)

var statusDesc = map[int]string{
	ErrLoginCodeInput:       "验证码错误，请重新输入",
	ErrLoginCodeRepeatSend:  "登陆验证码60s之内只能发一次",
	ErrLoginCodeSend:        "验证码发送失败，请稍后重试",
	ErrLoginCodeExpire:      "验证码已过期，请重新获取",
	ErrLoginNickName:        "操作失败请稍后再试",
}

func init() {
	//用户错误描述
	statusDesc[ErrUserNotExist] = "用户不存在"
	statusDesc[ErrUserWasDelete] = "用户已被删除"
	statusDesc[ErrUserUpdate] = "更新用户信息失败"
	statusDesc[ErrUserPhoneBinded] = "当前手机已经绑定了其他微信用户"
	statusDesc[ErrUserHasBindPhone] = "当前用户已经绑定手机号"
	statusDesc[ErrUserBindPhoneFailed] = "当前用户绑定手机号失败，请稍后重试"
	statusDesc[ErrUserHasNoPhone] = "当前用户未绑定手机号，请先绑定手机号"
	statusDesc[ErrUserPhoneIsNotNull] = "用户手机号不能为空"
}

func SResultError(code int, data interface{}) error {
	return NewServiceError(code, statusDesc[code], data, nil)
}

func SResultErrorC(code int, data interface{}, cause error) error {
	return NewServiceError(code, statusDesc[code], data, cause)
}

func SResultErrorMsg(code int, message string) error {
	return newServiceErrorNOData(code, message, nil)
}

func SResultErrorMsgC(code int, message string, cause error) error {
	return newServiceErrorNOData(code, message, cause)
}
