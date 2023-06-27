package routes

import (
	"BNMO/controllers"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func AccountVerifRoutes(route *gin.Engine) {
	auth := route.Group("/account-verif").Use(middleware.AdminMiddleware())
	auth.GET("/get", controllers.GetPendingAccounts)
	auth.PUT("/validate/:id/:status", controllers.ValidateAccount)
}
