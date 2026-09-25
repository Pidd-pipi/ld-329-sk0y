package controller

import (
	"net/http"
	"strconv"

	"cyskillswap/internal/errors"
	"cyskillswap/internal/middleware"
	"cyskillswap/internal/service"
	"github.com/gin-gonic/gin"
)

// Needs 返回带当前用户响应/预约状态的需求卡片列表。
func Needs(c *gin.Context) {
	c.JSON(http.StatusOK, service.ListNeedCards(middleware.CurrentActor(c)))
}

// Respond 提交响应（交换说明 + 空闲时段）。
func Respond(c *gin.Context) {
	needID, ok := parseID(c, "id")
	if !ok {
		return
	}
	var req service.RespondRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, errors.New(http.StatusBadRequest, "VALIDATION_FAILED", "请求内容格式不正确"))
		return
	}
	card, err := service.Respond(middleware.CurrentActor(c), needID, req)
	if err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusCreated, card)
}

// SelectCandidate 发起人从候选人中选定一人。
func SelectCandidate(c *gin.Context) {
	needID, ok := parseID(c, "id")
	if !ok {
		return
	}
	responseID, ok := parseID(c, "responseId")
	if !ok {
		return
	}
	apt, err := service.Select(middleware.CurrentActor(c), needID, responseID)
	if err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusCreated, apt)
}

func parseID(c *gin.Context, key string) (int, bool) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil || id <= 0 {
		writeError(c, errors.New(http.StatusBadRequest, "VALIDATION_FAILED", "无效的资源编号"))
		return 0, false
	}
	return id, true
}
