package errors

import "fmt"

// 业务错误码，前后端共用约定
const (
	CodeValidation        = "VALIDATION_FAILED"
	CodeNeedNotFound      = "NEED_NOT_FOUND"
	CodeNeedClosed        = "NEED_CLOSED"
	CodeResponseNotFound  = "RESPONSE_NOT_FOUND"
	CodeDuplicateResponse = "DUPLICATE_RESPONSE"
	CodeForbidden         = "FORBIDDEN"
	CodeScheduleConflict  = "SCHEDULE_CONFLICT"
)

func Validation(message string) BusinessError { return New(CodeValidation, message) }

func NeedNotFound() BusinessError { return New(CodeNeedNotFound, "需求不存在") }

func NeedClosed() BusinessError {
	return New(CodeNeedClosed, "该需求已定人，停止接收新的响应")
}

func ResponseNotFound() BusinessError {
	return New(CodeResponseNotFound, "响应不存在或不属于该需求")
}

func DuplicateResponse() BusinessError {
	return New(CodeDuplicateResponse, "你已响应过该需求，请勿重复报名")
}

func Forbidden(message string) BusinessError { return New(CodeForbidden, message) }

// ScheduleConflict 被选同学在同一时段已有其他预约
func ScheduleConflict(student, slot, detail string) BusinessError {
	return New(CodeScheduleConflict,
		fmt.Sprintf("%s 在 %s 已有其他预约（%s），请选择其他候选人", student, slot, detail))
}
