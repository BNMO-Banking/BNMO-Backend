package main

import (
	"BNMO/database"
	"BNMO/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"math/rand"
)

var (
	router = gin.Default()
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file", err.Error())
	}

	log.Println("Env successfully loaded")

	// Initialize database using GORM
	database.Initialize()

	// Setup random seed
	rand.Seed(time.Now().UnixNano())

	// Set up CORS policy
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "Static", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.AuthRoutes(router)
	routes.ProfileRoutes(router)
	routes.AccountVerifRoutes(router)
	routes.RequestRoutes(router)
	routes.RequestVerifRoutes(router)
	routes.PinRoutes(router)
	routes.AssociateRoutes(router)
	routes.TransferRoutes(router)
	router.Run()
}
