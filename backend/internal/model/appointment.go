package model

// Appointment 交换预约，Slot 用于同一时段的冲突检测
type Appointment struct {
	ID        int    `json:"id"`
	NeedID    int    `json:"needId"`
	Pair      string `json:"pair"`
	Requester string `json:"requester"`
	Provider  string `json:"provider"`
	Time      string `json:"time"`
	Slot      string `json:"slot"`
	Place     string `json:"place"`
	Status    string `json:"status"`
	Agenda    string `json:"agenda"`
}
