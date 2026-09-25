package repository

import (
	"sync"
	"time"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

// Store 是进程内演示存储。真实环境应替换为 MySQL 实现，service 层接口保持不变。
type Store struct {
	mu sync.RWMutex

	skills       []model.Skill
	needs        []model.Need
	responses    []model.Response
	matches      []model.Match
	appointments []model.Appointment
	reviews      []model.Review
	messages     []model.Conversation
	profile      model.Profile

	nextResponseID    int
	nextAppointmentID int
}

var seedTime = time.Date(2026, 9, 24, 20, 0, 0, 0, time.Local)

// NewStore 构造带有演示数据的存储。
func NewStore() *Store {
	s := &Store{nextResponseID: 10, nextAppointmentID: 10}
	s.seed()
	return s
}

// Lock / RLock 暴露给 service 层做跨实体事务性操作。
func (s *Store) Lock()    { s.mu.Lock() }
func (s *Store) Unlock()  { s.mu.Unlock() }
func (s *Store) RLock()   { s.mu.RLock() }
func (s *Store) RUnlock() { s.mu.RUnlock() }

func (s *Store) seed() {
	s.skills = []model.Skill{
		{ID: 1, Owner: "林澈", Title: "毕业照人像摄影", Category: "摄影", Level: 92, Campus: "东校区", Description: "提供构图、修图和毕业季跟拍，可交换吉他入门课。", TimeSlots: []string{"周三晚", "周六上午"}, Rewards: []string{"技能交换", "请吃饭"}, Portfolio: "12组校园人像作品"},
		{ID: 2, Owner: "周芮", Title: "Python 数据分析", Category: "编程", Level: 88, Campus: "中心校区", Description: "pandas、可视化、论文数据清洗辅导，接受小额报酬。", TimeSlots: []string{"周二晚", "周日全天"}, Rewards: []string{"技能交换", "小额报酬"}, Portfolio: "3份课程项目证书"},
		{ID: 3, Owner: "孟野", Title: "民谣吉他陪练", Category: "乐器", Level: 81, Campus: "西校区", Description: "节奏型、弹唱和舞台经验分享，想找人拍宣传照。", TimeSlots: []string{"周三晚", "周六上午"}, Rewards: []string{"技能交换", "无偿"}, Portfolio: "校园音乐节演出视频"},
	}

	s.needs = []model.Need{
		{ID: 1, Requester: "孟野", Title: "找人帮忙拍乐队宣传照", Category: "摄影", Campus: "西校区", ExpectTime: constants.SlotLabel(constants.SlotSatMorning), SlotCode: constants.SlotSatMorning, BudgetType: "技能交换", Description: "可交换 3 次吉他课，希望会调色和室外构图。", Status: constants.NeedStatusOpen},
		{ID: 2, Requester: "许安", Title: "求教 Python 数据分析", Category: "编程", Campus: "中心校区", ExpectTime: constants.SlotLabel(constants.SlotTueEvening), SlotCode: constants.SlotTueEvening, BudgetType: "小额报酬", Description: "论文问卷数据需要清洗和画图，最好有 pandas 经验。", Status: constants.NeedStatusOpen},
		{ID: 3, Requester: "林澈", Title: "想学吉他扫弦入门", Category: "乐器", Campus: "东校区", ExpectTime: constants.SlotLabel(constants.SlotWedEvening), SlotCode: constants.SlotWedEvening, BudgetType: "技能交换", Description: "用摄影课交换吉他基础，希望同校区或线上。", Status: constants.NeedStatusOpen},
	}

	// 既有预约：林澈周六上午已与顾远有约——选林澈拍宣传照时会命中冲突。
	s.appointments = []model.Appointment{
		{ID: 1, NeedID: 0, Requester: "顾远", Respondent: "林澈", SlotCode: constants.SlotSatMorning, Time: constants.SlotLabel(constants.SlotSatMorning) + " 10:00", Place: "东校区湖边", Agenda: "社团招新跟拍，约两小时", Status: constants.AppointmentConfirmed, RequesterOK: true, RespondentOK: true, CreatedAt: seedTime},
		{ID: 2, NeedID: 0, Requester: "许安", Respondent: "周芮", SlotCode: constants.SlotTueEvening, Time: constants.SlotLabel(constants.SlotTueEvening) + " 19:30", Place: "线上会议室", Agenda: "导入问卷 CSV 并完成基础可视化", Status: constants.AppointmentPending, RequesterOK: true, RespondentOK: false, CreatedAt: seedTime},
	}

	s.responses = []model.Response{
		{ID: 1, NeedID: 1, Respondent: "林澈", OfferNote: "擅长室外自然光和调色，可用毕业照跟拍经验交换吉他课。", FreeSlots: []string{constants.SlotWedEvening, constants.SlotSatMorning}, Status: constants.ResponseWaiting, CreatedAt: seedTime},
		{ID: 2, NeedID: 1, Respondent: "唐宁", OfferNote: "摄影社成员，带过乐队现场拍摄，设备齐全可出原片。", FreeSlots: []string{constants.SlotSatMorning, constants.SlotSunAllDay}, Status: constants.ResponseWaiting, CreatedAt: seedTime},
		{ID: 3, NeedID: 1, Respondent: "韩梅", OfferNote: "会后期精修，可先看作品集再决定，接受技能交换。", FreeSlots: []string{constants.SlotFriEvening, constants.SlotSatMorning}, Status: constants.ResponseWaiting, CreatedAt: seedTime},
		{ID: 4, NeedID: 2, Respondent: "周芮", OfferNote: "pandas 与论文图表经验丰富，可按次结算报酬。", FreeSlots: []string{constants.SlotTueEvening, constants.SlotSunAllDay}, Status: constants.ResponseWaiting, CreatedAt: seedTime},
		{ID: 5, NeedID: 2, Respondent: "唐宁", OfferNote: "统计方向研究生，熟悉问卷信效度和回归分析。", FreeSlots: []string{constants.SlotThuEvening, constants.SlotSunAllDay}, Status: constants.ResponseWaiting, CreatedAt: seedTime},
		{ID: 6, NeedID: 3, Respondent: "孟野", OfferNote: "可从扫弦节奏型讲起，顺带教弹唱配合，想换一次人像拍摄。", FreeSlots: []string{constants.SlotWedEvening, constants.SlotSatMorning}, Status: constants.ResponseWaiting, CreatedAt: seedTime},
		{ID: 7, NeedID: 3, Respondent: "周芮", OfferNote: "民谣弹唱业余三年，适合陪练入门，时间固定周三晚。", FreeSlots: []string{constants.SlotWedEvening}, Status: constants.ResponseWaiting, CreatedAt: seedTime},
	}

	s.matches = []model.Match{
		{ID: 1, Provider: "林澈", Learner: "孟野", OfferSkill: "毕业照人像摄影", WantedSkill: "民谣吉他陪练", Score: 96, CommonSlots: []string{"周三晚", "周六上午"}, Recommendation: "互补技能明确，双方均接受技能交换。"},
		{ID: 2, Provider: "周芮", Learner: "许安", OfferSkill: "Python 数据分析", WantedSkill: "论文数据清洗", Score: 89, CommonSlots: []string{"周二晚"}, Recommendation: "时间匹配且需求描述命中 pandas/可视化。"},
		{ID: 3, Provider: "孟野", Learner: "林澈", OfferSkill: "民谣吉他陪练", WantedSkill: "宣传照拍摄", Score: 91, CommonSlots: []string{"周六上午"}, Recommendation: "互换回报类型一致，信用分权重较高。"},
	}

	s.reviews = []model.Review{
		{ID: 1, From: "孟野", To: "林澈", Rating: 5, Content: "构图建议很细，成片当天就给了预览。"},
		{ID: 2, From: "林澈", To: "孟野", Rating: 5, Content: "吉他入门节奏拆得很清楚，课后还发了练习谱。"},
	}

	s.messages = []model.Conversation{
		{ID: 1, WithUser: "孟野", Unread: 2, Messages: []string{"周六湖边光线不错", "我带两套衣服可以吗？"}},
		{ID: 2, WithUser: "系统通知", Unread: 1, Messages: []string{"你与周芮的 Python 数据分析预约待确认。"}},
	}

	s.profile = model.Profile{
		Name: "林澈", Major: "新闻传播 2023", CreditScore: 91, CreditLevel: constants.CreditGold,
		SkillWall: s.skills[:1],
		Radar:     map[string]int{"摄影": 92, "修图": 86, "沟通": 90, "编程": 42, "乐器": 35},
		History:   []string{"完成毕业照拍摄交换", "响应 Python 数据分析需求", "预约吉他入门课"},
		Reviews:   s.reviews,
	}
}
