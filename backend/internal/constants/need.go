package constants

// 需求状态
const (
	NeedStatusOpen    = "招募中"
	NeedStatusMatched = "已定人"
)

// 响应状态
const (
	ResponseStatusPending  = "待选定"
	ResponseStatusSelected = "已入选"
	ResponseStatusRejected = "未入选"
)

// 预约状态
const (
	AppointmentStatusPending   = "待双方确认"
	AppointmentStatusConfirmed = "双方已确认"
)

// TimeSlots 可选择的空闲时段
var TimeSlots = []string{"周一晚", "周二晚", "周三晚", "周四晚", "周五晚", "周六上午", "周六下午", "周日全天"}

// DefaultAppointmentPlace 预约生成时地点待双方协商
const DefaultAppointmentPlace = "待双方协商"
