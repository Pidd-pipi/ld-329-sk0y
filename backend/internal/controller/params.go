package controller

import (
	"strconv"

	"cyskillswap/internal/errors"
	"github.com/gin-gonic/gin"
)

// parseNeedID 解析路径中的需求 ID
func parseNeedID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		respondError(c, errors.Validation("无效的需求 ID"))
		return 0, false
	}
	return id, true
}
