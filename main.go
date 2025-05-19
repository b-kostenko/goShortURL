package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goShortURL/src/auth"
	"goShortURL/src/database"
	"goShortURL/src/shortlink"
	"time"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "OK", "code": 200, "time": time.Now()})
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	database.ConnectDatabase()
	auth.SyncModels()
	shortlink.SyncModels()
}

func main() {
	router := gin.Default()
	router.GET("/health-check", healthCheck)
	auth.Routers(router)
	shortlink.Routers(router)

	err := router.Run("localhost:8081")
	if err != nil {
		panic(err)
	}
}
