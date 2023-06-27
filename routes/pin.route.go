package routes

import (
	"BNMO/controllers"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func PinRoutes(route *gin.Engine) {
	pin := route.Group("/pin").Use(middleware.CustomerMiddleware())
	pin.POST("/set", controllers.SetPin)
	pin.POST("/compare", controllers.ComparePin)
}
