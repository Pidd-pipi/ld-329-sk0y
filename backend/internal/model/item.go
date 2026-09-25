package model

import "time"

type Skill struct {
	ID          int      `json:"id"`
	Owner       string   `json:"owner"`
	Title       string   `json:"title"`
	Category    string   `json:"category"`
	Level       int      `json:"level"`
	Campus      string   `json:"campus"`
	Description string   `json:"description"`
	TimeSlots   []string `json:"timeSlots"`
	Rewards     []string `json:"rewards"`
	Portfolio   string   `json:"portfolio"`
}

type Need struct {
	ID          int    `json:"id"`
	Requester   string `json:"requester"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Campus      string `json:"campus"`
	ExpectTime  string `json:"expectTime"`
	SlotCode    string `json:"slotCode"`
	BudgetType  string `json:"budgetType"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// Response 是同学对某条需求提交的报名（交换说明 + 空闲时段）。
type Response struct {
	ID         int       `json:"id"`
	NeedID     int       `json:"needId"`
	Respondent string    `json:"respondent"`
	OfferNote  string    `json:"offerNote"`
	FreeSlots  []string  `json:"freeSlots"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
}

type Match struct {
	ID             int      `json:"id"`
	Provider       string   `json:"provider"`
	Learner        string   `json:"learner"`
	OfferSkill     string   `json:"offerSkill"`
	WantedSkill    string   `json:"wantedSkill"`
	Score          int      `json:"score"`
	CommonSlots    []string `json:"commonSlots"`
	Recommendation string   `json:"recommendation"`
}

// Appointment 是发起人选定候选人后按需求时间生成的预约，需双方确认生效。
type Appointment struct {
	ID           int       `json:"id"`
	NeedID       int       `json:"needId"`
	Requester    string    `json:"requester"`
	Respondent   string    `json:"respondent"`
	SlotCode     string    `json:"slotCode"`
	Time         string    `json:"time"`
	Place        string    `json:"place"`
	Agenda       string    `json:"agenda"`
	Status       string    `json:"status"`
	RequesterOK  bool      `json:"requesterConfirmed"`
	RespondentOK bool      `json:"respondentConfirmed"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Review struct {
	ID      int    `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Rating  int    `json:"rating"`
	Content string `json:"content"`
}

type Conversation struct {
	ID       int      `json:"id"`
	WithUser string   `json:"withUser"`
	Unread   int      `json:"unread"`
	Messages []string `json:"messages"`
}

type Profile struct {
	Name        string         `json:"name"`
	Major       string         `json:"major"`
	CreditScore int            `json:"creditScore"`
	CreditLevel string         `json:"creditLevel"`
	SkillWall   []Skill        `json:"skillWall"`
	Radar       map[string]int `json:"radar"`
	History     []string       `json:"history"`
	Reviews     []Review       `json:"reviews"`
}

type Overview struct {
	Service         string         `json:"service"`
	CurrentUser     string         `json:"currentUser"`
	SwitchableUsers []string       `json:"switchableUsers"`
	Categories      []string       `json:"categories"`
	Slots           []SlotOption   `json:"slots"`
	Metrics         map[string]int `json:"metrics"`
	Skills          []Skill        `json:"skills"`
	Needs           []NeedCard     `json:"needs"`
	Matches         []Match        `json:"matches"`
	Appointments    []Appointment  `json:"appointments"`
	Reviews         []Review       `json:"reviews"`
	Messages        []Conversation `json:"messages"`
	Profile         Profile        `json:"profile"`
}

// SlotOption 暴露给前端的空闲时段字典。
type SlotOption struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}
