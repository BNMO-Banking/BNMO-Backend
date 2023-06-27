package routes

import (
	"BNMO/controllers"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func AccountVerifRoutes(route *gin.Engine) {
	account_verif := route.Group("/account-verif").Use(middleware.AdminMiddleware())
	account_verif.GET("/get", controllers.GetPendingAccounts)
	account_verif.PUT("/validate/:id/:status", controllers.ValidateAccount)
}
