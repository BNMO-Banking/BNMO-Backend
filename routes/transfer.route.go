package routes

import (
	"BNMO/controllers"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func TransferRoutes(route *gin.Engine) {
	transfer := route.Group("/transfer").Use(middleware.CustomerMiddleware())
	transfer.POST("/send", controllers.Transfer)
}
