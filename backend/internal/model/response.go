package model

// NeedResponse 同学对某条需求提交的响应（报名）
type NeedResponse struct {
	ID      int      `json:"id"`
	NeedID  int      `json:"needId"`
	Student string   `json:"student"`
	Note    string   `json:"note"`
	Slots   []string `json:"slots"`
	Status  string   `json:"status"`
}

// ResponseInput 提交响应的请求体：交换说明 + 空闲时段
type ResponseInput struct {
	Note  string   `json:"note"`
	Slots []string `json:"slots"`
}

// SelectInput 发起人选定候选人的请求体
type SelectInput struct {
	ResponseID int `json:"responseId"`
}
