package routes

import (
	"BNMO/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.POST("/register", controllers.RegisterAccount)
	auth.POST("/login", controllers.LoginAccount)
	auth.POST("/logout", controllers.LogoutAccount)
}
