package controller

import (
	"net/http"

	"cyskillswap/internal/errors"
	"github.com/gin-gonic/gin"
)

// respondError 将业务异常映射为统一的错误响应
func respondError(c *gin.Context, err error) {
	if biz, ok := err.(errors.BusinessError); ok {
		c.JSON(statusFor(biz.Code), biz)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL", "message": "服务器内部错误"})
}

func statusFor(code string) int {
	switch code {
	case errors.CodeNeedNotFound, errors.CodeResponseNotFound:
		return http.StatusNotFound
	case errors.CodeForbidden:
		return http.StatusForbidden
	case errors.CodeNeedClosed, errors.CodeDuplicateResponse, errors.CodeScheduleConflict:
		return http.StatusConflict
	default:
		return http.StatusBadRequest
	}
}
