package routes

import (
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func FilesRoutes(route *gin.Engine) {
	route.Static("/images", "images").Use(middleware.ProtectedMiddleware())
}
