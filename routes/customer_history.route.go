package routes

import (
	"BNMO/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerHistoryRoutes(route *gin.Engine) {
	customer_history := route.Group("/customer-history")
	customer_history.GET("/get-request/:id", controllers.GetRequestHistory)
	customer_history.GET("/get-transfer/:id", controllers.GetTransferHistory)
}
