package cerrors

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
)

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
	ErrLoginErr 				  = 10002
)

var codeMessage = map[int]string{
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
	ErrLoginErr:				   "不存在的用户或错误的密码",
}

type ServiceError interface {
	Error() string
	ErrCode() int
	ErrMessage() string
	ErrData() interface{}
}

type serviceError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	cause   error
}

func (s *serviceError) Error() string {
	errBuf := bytes.NewBufferString("")
	errBuf.WriteString(fmt.Sprintf("%d%s \n", s.Code, s.Message))
	if s.cause != nil {
		errBuf.WriteString("\n")
		errBuf.WriteString(fmt.Sprintf("%+v", s.cause))
	}

	return errBuf.String()
}

func (s *serviceError) ErrCode() int {
	return s.Code
}

func (s *serviceError) ErrMessage() string {
	return s.Message
}

func (s *serviceError) ErrData() interface{} {
	return s.Data
}

func NewServiceError(code int, message string, data interface{}, cause error) error {
	if data == nil {
		data = make(map[string]interface{})
	}

	if cause != nil {
		cause = errors.WithStack(cause)
	}

	return &serviceError{
		Code:    code,
		Message: message,
		Data:    data,
		cause:   cause,
	}
}

func NewServiceErrorCode(code int, message string) error {
	if message == "" {
		message = codeMessage[code]
	}

	if message == "" {
		message = statusDesc[code]
	}

	return newServiceErrorNOData(code, message, nil)
}

func NewServiceErrorCodeC(code int, message string, cause error) error {
	if message == "" {
		message = codeMessage[code]
	}

	if message == "" {
		message = statusDesc[code]
	}

	return newServiceErrorNOData(code, message, cause)
}

func New404ServiceError(message string) error {
	return newServiceErrorNOData(404, message, nil)
}

func New400ServiceError(message string) error {
	if message == "" {
		message = "出现意外，填写格式问题"
	}
	return newServiceErrorNOData(400, message, nil)
}

func New400ServiceErrorC(message string, cause error) error {
	if message == "" {
		message = "出现意外，填写格式问题"
	}
	return newServiceErrorNOData(400, message, cause)
}

func New500ServiceError() error {
	return newServiceErrorNOData(500, "服务器内部错误，请稍后重试", nil)
}

func New600ServiceError() error {
	return newServiceErrorNOData(600, "出现意外，请稍后重试", nil)
}

func New600ServiceErrorC(cause error) error {
	return newServiceErrorNOData(600, "出现意外，请稍后重试", cause)
}

func New700ServiceError() error {
	return newServiceErrorNOData(700, "出现意外，请稍后重试", nil)
}

func New700ServiceErrorC(cause error) error {
	return newServiceErrorNOData(700, "出现意外，请稍后重试", cause)
}

func newServiceErrorNOData(code int, message string, cause error) error {
	return NewServiceError(code, message, nil, cause)
}


