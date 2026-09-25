package controller

import (
	"net/http"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/service"
	"github.com/gin-gonic/gin"
)

// Needs 需求列表（含当前用户的响应与预约状态）
func Needs(c *gin.Context) {
	c.JSON(http.StatusOK, service.Needs())
}

// Responses 发起人查看某条需求的候选人列表
func Responses(c *gin.Context) {
	needID, ok := parseNeedID(c)
	if !ok {
		return
	}
	candidates, err := service.Candidates(constants.CurrentUser, needID)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, candidates)
}

// Respond 同学提交响应：交换说明 + 空闲时段
func Respond(c *gin.Context) {
	needID, ok := parseNeedID(c)
	if !ok {
		return
	}
	var input model.ResponseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondError(c, errors.Validation("请求格式不正确"))
		return
	}
	resp, err := service.SubmitResponse(constants.CurrentUser, needID, input)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// Select 发起人选定候选人，生成待双方确认的预约
func Select(c *gin.Context) {
	needID, ok := parseNeedID(c)
	if !ok {
		return
	}
	var input model.SelectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondError(c, errors.Validation("请求格式不正确"))
		return
	}
	appt, err := service.SelectCandidate(constants.CurrentUser, needID, input.ResponseID)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, appt)
}
