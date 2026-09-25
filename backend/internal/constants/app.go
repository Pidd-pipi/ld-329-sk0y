package constants

const (
	ServiceName = "cyskillswap"
	APIPrefix   = "/api"

	// CurrentUser 是未接入登录时的默认当前用户；可通过请求头 X-User-Name 切换。
	CurrentUser = "林澈"
	// ActorHeader 允许前端在演示环境切换身份，登录接入后由 JWT 中间件替换。
	ActorHeader = "X-User-Name"
)

var SkillCategories = []string{"摄影", "编程", "乐器", "外语", "平面设计", "健身指导"}

// SwitchableUsers 是前端身份切换器可选的演示用户。
var SwitchableUsers = []string{"林澈", "孟野", "许安", "周芮", "唐宁", "韩梅"}

const (
	CreditBronze = "青铜互助者"
	CreditSilver = "白银协作者"
	CreditGold   = "黄金导师"
)
