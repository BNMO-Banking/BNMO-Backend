package routes

import (
	"BNMO/controllers"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func AssociateRoutes(route *gin.Engine) {
	associates := route.Group("/associates").Use(middleware.CustomerMiddleware())
	associates.POST("/add", controllers.AddAssociates)
	associates.GET("/check/:number", controllers.CheckAssociates)
	associates.GET("/get", controllers.GetAssociates)
}
