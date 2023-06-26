package routes

import (
	"BNMO/controllers"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(route *gin.Engine) {
	profile := route.Group("/profile").Use(middleware.CustomerMiddleware())
	profile.GET("/get/:id", controllers.GetProfile)
	profile.POST("/edit/:id", controllers.EditProfile)
}
