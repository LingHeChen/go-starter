package xerr

import "fmt"

type Error interface {
	error
	GetCode() int
	GetMsg() string
}

// CodeError 自定义业务错误
type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 实现 error 接口
func (e *CodeError) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

func (e *CodeError) GetCode() int {
	return e.Code
}

func (e *CodeError) GetMsg() string {
	return e.Msg
}

// New 创建一个自定义错误
func New(code int, msg string) error {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}

// NewErrMsg 创建一个通用错误（使用默认错误码 -1）
func NewErrMsg(msg string) error {
	return &CodeError{
		Code: ServerCommonError,
		Msg:  msg,
	}
}

// NewErrCode 仅通过错误码创建（Msg 使用该码对应的默认文案，需配合 map 使用，这里简化直接返回）
func NewErrCode(code int) error {
	return &CodeError{
		Code: code,
		Msg:  "Business Error", // 实际项目中这里应该查 Map 获取默认文案
	}
}
