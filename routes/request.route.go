package routes

import (
	"BNMO/controllers"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func RequestRoutes(route *gin.Engine) {
	request := route.Group("/request").Use(middleware.CustomerMiddleware())
	request.POST("/add", controllers.AddRequest)
}
