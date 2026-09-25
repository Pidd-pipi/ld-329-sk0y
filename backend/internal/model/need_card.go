package model

// NeedCard 是需求列表卡片视图：在需求基础上附带候选人、当前用户的响应与预约状态。
type NeedCard struct {
	Need
	Responses     []Response   `json:"responses"`
	ResponseCount int          `json:"responseCount"`
	MyResponse    *Response    `json:"myResponse,omitempty"`
	Appointment   *Appointment `json:"appointment,omitempty"`
}
