package controller

import (
	"cyskillswap/internal/errors"
	"github.com/gin-gonic/gin"
)

// writeError 把业务异常统一翻译成 HTTP 响应；未知错误按 500 处理。
func writeError(c *gin.Context, err error) {
	if biz, ok := err.(errors.BusinessError); ok {
		c.JSON(biz.Status, gin.H{
			"error":   biz.Code,
			"message": biz.Message,
			"details": biz.Details,
		})
		return
	}
	c.JSON(500, gin.H{"error": "INTERNAL_ERROR", "message": "服务暂时不可用"})
}
