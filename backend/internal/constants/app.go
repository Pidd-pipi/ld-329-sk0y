package constants

const (
	ServiceName = "cyskillswap"
	APIPrefix   = "/api"
)

// CurrentUser 演示环境的当前登录用户（正式环境由 JWT 解析得到）
const CurrentUser = "林澈"

var SkillCategories = []string{"摄影", "编程", "乐器", "外语", "平面设计", "健身指导"}

const (
	CreditBronze = "青铜互助者"
	CreditSilver = "白银协作者"
	CreditGold   = "黄金导师"
)
