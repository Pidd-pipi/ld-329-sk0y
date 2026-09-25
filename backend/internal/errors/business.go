package errors

// BusinessError 业务异常，统一携带错误码与对用户友好的提示
type BusinessError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e BusinessError) Error() string { return e.Message }

func New(code, message string) BusinessError {
	return BusinessError{Code: code, Message: message}
}
