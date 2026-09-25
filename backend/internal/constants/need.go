package constants

// 需求状态：open 招募中（接收响应）；booked 已选定候选人（停止接收响应）。
const (
	NeedStatusOpen   = "open"
	NeedStatusBooked = "booked"
)

// 响应状态：waiting 候选中；selected 已被发起人选中；released 需求选定他人后未中选。
const (
	ResponseWaiting  = "waiting"
	ResponseSelected = "selected"
	ResponseReleased = "released"
)

// 预约状态：pending 待双方确认；confirmed 双方已确认；cancelled 已取消（发起人改选）。
const (
	AppointmentPending   = "pending"
	AppointmentConfirmed = "confirmed"
	AppointmentCancelled = "cancelled"
)
