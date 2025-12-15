package xerr

// 常用通用错误码
const (
	OK                = 0
	ServerCommonError = -1 // 或者 500，看你喜好
	RequestParamError = 400
	TokenExpired      = 401
)

// 业务错误码示例 (建议后续按模块划分，如 User 模块 10001)
const (
	UserNotFound = 10001
	UserPwdError = 10002
)
