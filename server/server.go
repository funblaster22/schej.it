package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"schej.it/server/db"
	"schej.it/server/routes"
)

func main() {
	router := gin.Default()

	// Load .env variables
	loadDotEnv()

	// Cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Init database
	closeConnection := db.Init()
	defer closeConnection()

	// Session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	// Init routes
	routes.InitAuth(router)
	routes.InitUser(router)
	routes.InitEvents(router)

	// Run server
	router.Run(":3000")
}

// Load .env variables
func loadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
}
