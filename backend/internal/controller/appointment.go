package controller

import (
	"net/http"

	"cyskillswap/internal/middleware"
	"cyskillswap/internal/service"
	"github.com/gin-gonic/gin"
)

func Appointments(c *gin.Context) {
	c.JSON(http.StatusOK, service.ListAppointments())
}

// ConfirmAppointment 预约双方点击确认。
func ConfirmAppointment(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	apt, err := service.ConfirmAppointment(middleware.CurrentActor(c), id)
	if err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusOK, apt)
}

// CancelAppointment 取消待确认预约并释放需求（可改选别人）。
func CancelAppointment(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	apt, err := service.CancelAppointment(middleware.CurrentActor(c), id)
	if err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusOK, apt)
}
