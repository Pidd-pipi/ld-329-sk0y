package constants

// 集中维护业务错误码，禁止在业务代码里抛裸字符串。
const (
	CodeValidationFailed    = "VALIDATION_FAILED"
	CodeNeedNotFound        = "NEED_NOT_FOUND"
	CodeNeedClosed          = "NEED_CLOSED"
	CodeNeedAlreadyBooked   = "NEED_ALREADY_BOOKED"
	CodeOwnNeed             = "CANNOT_RESPOND_OWN_NEED"
	CodeDuplicateResponse   = "NEED_ALREADY_RESPONDED"
	CodeResponseNotFound    = "RESPONSE_NOT_FOUND"
	CodeSlotMismatch        = "RESPONDENT_SLOT_UNAVAILABLE"
	CodeNotRequester        = "NOT_NEED_REQUESTER"
	CodeRespondentBusy      = "RESPONDENT_SLOT_CONFLICT"
	CodeApptNotFound        = "APPOINTMENT_NOT_FOUND"
	CodeNotAppointmentParty = "NOT_APPOINTMENT_PARTY"
	CodeAlreadyConfirmed    = "APPOINTMENT_ALREADY_CONFIRMED"
	CodeAppointmentClosed   = "APPOINTMENT_CLOSED"
	CodeAppointmentLocked   = "APPOINTMENT_NOT_PENDING"
)
