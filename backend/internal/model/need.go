package model

// Need 学生发布的技能求助需求
type Need struct {
	ID          int    `json:"id"`
	Requester   string `json:"requester"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Campus      string `json:"campus"`
	ExpectTime  string `json:"expectTime"`
	Slot        string `json:"slot"`
	BudgetType  string `json:"budgetType"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Responses   int    `json:"responses"`
}

// NeedView 需求卡片视图，附带当前用户视角的响应、预约与候选人信息
type NeedView struct {
	Need
	MyResponse    *NeedResponse  `json:"myResponse,omitempty"`
	MyAppointment *Appointment   `json:"myAppointment,omitempty"`
	Candidates    []NeedResponse `json:"candidates,omitempty"`
}
