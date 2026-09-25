package errors

import "cyskillswap/internal/constants"

// BusinessError 是带错误码的业务异常，由 controller 统一转成 HTTP 响应。
type BusinessError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
	// Details 携带冲突详情等结构化信息，前端可直接展示。
	Details map[string]any `json:"details,omitempty"`
}

func (e BusinessError) Error() string { return e.Message }

func New(status int, code, message string) BusinessError {
	return BusinessError{Code: code, Message: message, Status: status}
}

func NewWithDetails(status int, code, message string, details map[string]any) BusinessError {
	return BusinessError{Code: code, Message: message, Status: status, Details: details}
}

// 预置业务异常，消息文案集中管理。
var (
	ErrNeedNotFound      = New(404, constants.CodeNeedNotFound, "需求不存在或已下线")
	ErrResponseNotFound  = New(404, constants.CodeResponseNotFound, "报名记录不存在")
	ErrApptNotFound      = New(404, constants.CodeApptNotFound, "预约不存在")
	ErrNeedClosed        = New(409, constants.CodeNeedClosed, "该需求已选定同学，不再接收新的响应")
	ErrOwnNeed           = New(400, constants.CodeOwnNeed, "不能响应自己发布的需求")
	ErrDuplicateResponse = New(409, constants.CodeDuplicateResponse, "你已响应过这条需求，请勿重复报名")
	ErrNotRequester      = New(403, constants.CodeNotRequester, "只有需求发起人可以选定候选人")
	ErrNotApptParty      = New(403, constants.CodeNotAppointmentParty, "只有预约双方可以操作该预约")
	ErrAlreadyConfirmed  = New(409, constants.CodeAlreadyConfirmed, "该预约已双方确认，无需重复确认")
	ErrAppointmentClosed = New(409, constants.CodeAppointmentClosed, "该预约已取消，无法继续操作")
)
