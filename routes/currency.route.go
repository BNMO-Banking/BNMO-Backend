package routes

import (
	"BNMO/controllers"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func CurrencyRoutes(route *gin.Engine) {
	currency := route.Group("/currency").Use(middleware.CustomerMiddleware())
	currency.GET("/get-symbols", controllers.GetSymbols)
}
