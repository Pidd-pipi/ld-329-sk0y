package routes

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/controller"
	"cyskillswap/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	api := r.Group(constants.APIPrefix)
	api.Use(middleware.Actor)

	api.GET("/health", controller.Health)
	api.GET("/dashboard/overview", controller.Overview)

	api.GET("/skills", controller.Skills)

	// 需求、响应与选定
	api.GET("/needs", controller.Needs)
	api.POST("/needs/:id/responses", controller.Respond)
	api.POST("/needs/:id/responses/:responseId/select", controller.SelectCandidate)

	api.GET("/matches", controller.Matches)

	// 预约确认与取消改选
	api.GET("/appointments", controller.Appointments)
	api.POST("/appointments/:id/confirm", controller.ConfirmAppointment)
	api.POST("/appointments/:id/cancel", controller.CancelAppointment)

	api.GET("/reviews", controller.Reviews)
	api.GET("/messages", controller.Messages)
	api.GET("/profile", controller.Profile)
}
