package middleware

import (
	"cyskillswap/internal/constants"
	"github.com/gin-gonic/gin"
)

// ActorContextKey 是当前登录同学在 gin.Context 中的键。
const ActorContextKey = "actor"

// Actor 在 JWT 接入前解析当前同学身份：优先读 X-User-Name 头，回退为演示默认用户。
func Actor(c *gin.Context) {
	actor := c.GetHeader(constants.ActorHeader)
	if actor == "" {
		actor = constants.CurrentUser
	}
	c.Set(ActorContextKey, actor)
	c.Next()
}

// CurrentActor 从请求上下文中取出当前同学姓名。
func CurrentActor(c *gin.Context) string {
	if v, ok := c.Get(ActorContextKey); ok {
		if name, ok := v.(string); ok && name != "" {
			return name
		}
	}
	return constants.CurrentUser
}
