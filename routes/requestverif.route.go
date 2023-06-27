package routes

import (
	"BNMO/controllers"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func RequestVerifRoutes(route *gin.Engine) {
	request_verif := route.Group("/request-verif").Use(middleware.AdminMiddleware())
	request_verif.GET("/get", controllers.GetPendingRequests)
	request_verif.PUT("/validate/:id/:status", controllers.ValidateRequest)
}
