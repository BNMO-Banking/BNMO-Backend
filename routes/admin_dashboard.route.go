package routes

import (
	"BNMO/controllers"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func AdminDashboardRoutes(route *gin.Engine) {
	admin_dashboard := route.Group("/admin-dashboard").Use(middleware.AdminMiddleware())
	admin_dashboard.GET("/get-pending-list", controllers.GetPendingLists)
	admin_dashboard.GET("/get-account-statistics", controllers.GetNewAccountStatistics)
	admin_dashboard.GET("/get-request-statistics", controllers.GetRequestTypeStatistics)
	admin_dashboard.GET("/get-bank-statistics", controllers.GetBankStatistics)
}
