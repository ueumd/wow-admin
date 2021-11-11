package constant
const (
	ErrDefaultServer         = 500
	ErrDefaultSql            = 600
	ErrDefaultRedis          = 700
	ErrDefaultOss            = 800
	ErrDefaultFileExceedSize = 801
	ErrDefaultParam          = 400
	ErrDefaultNotFound       = 404
	ErrDefaultNoPermission   = 401

	ErrDefaultToken               = 10000 //Token 错误
	ErrDefaultNeedLogin           = 10001 // 无Token 要求登陆
	ErrDefaultServerBusy          = 10010 // 服务器繁忙请稍后重试
	ErrDefaultTimeOut             = 10013 //操作超时
	ErrDefaultCheckCodeRepeatSend = 10014
	ErrDefaultCheckCodeSend       = 10015 //验证码发送失败
	ErrDefaultCheckCodeExpire     = 10016 //验证码已过期
	ErrDefaultCheckCodeInput      = 10017 //验证码输入错误
	ErrDefaultCheckCodeOldFailed  = 10018 //旧手机号未经过验证
)

var CodeMessage = map[int]string{
	ErrDefaultServer:              "服务器内部错误",
	ErrDefaultSql:                 "S数据访问出错",
	ErrDefaultRedis:               "R数据访问出错",
	ErrDefaultOss:                 "OSS数据访问出错",
	ErrDefaultFileExceedSize:      "上传文件大小超出限制10M",
	ErrDefaultParam:               "请求出现意外，填写格式问题",
	ErrDefaultNotFound:            "请求资源未找到",
	ErrDefaultNoPermission:        "当前用户没有操作权限",
	ErrDefaultTimeOut:             "操作超时",
	ErrDefaultCheckCodeExpire:     "验证码不正确或已失效",
	ErrDefaultCheckCodeSend:       "验证码发送失败",
	ErrDefaultCheckCodeInput:      "验证码不正确",
	ErrDefaultCheckCodeRepeatSend: "60s内重复请求验证码",
	ErrDefaultCheckCodeOldFailed:  "旧手机号未经过验证",
}
