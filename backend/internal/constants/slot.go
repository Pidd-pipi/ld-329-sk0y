package constants

// Slot 是可预约空闲时段的统一编码，避免业务代码里硬编码中文时段字符串。
type Slot struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

const (
	SlotTueEvening = "tue-evening"
	SlotWedEvening = "wed-evening"
	SlotThuEvening = "thu-evening"
	SlotFriEvening = "fri-evening"
	SlotSatMorning = "sat-morning"
	SlotSunAllDay  = "sun-all-day"
)

// FreeSlots 是报名时可选择的全部空闲时段（同时用于冲突检测的时间维度）。
var FreeSlots = []Slot{
	{Code: SlotTueEvening, Label: "周二晚"},
	{Code: SlotWedEvening, Label: "周三晚"},
	{Code: SlotThuEvening, Label: "周四晚"},
	{Code: SlotFriEvening, Label: "周五晚"},
	{Code: SlotSatMorning, Label: "周六上午"},
	{Code: SlotSunAllDay, Label: "周日全天"},
}

// SlotLabel 返回时段的中文展示名，未知编码回退为编码本身。
func SlotLabel(code string) string {
	for _, slot := range FreeSlots {
		if slot.Code == code {
			return slot.Label
		}
	}
	return code
}

// IsValidSlot 判断时段编码是否在系统登记范围内。
func IsValidSlot(code string) bool {
	for _, slot := range FreeSlots {
		if slot.Code == code {
			return true
		}
	}
	return false
}
