package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goShortURL/internal/api/handlers"
	"goShortURL/internal/api/middleware"
	"goShortURL/internal/database/postgres"
	"goShortURL/internal/models"
	"goShortURL/internal/repository"
	"goShortURL/internal/services"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

}

func main() {
	db, err := postgres.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}
	router := gin.Default()
	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	handler := handlers.NewHandler(authService, userRepo)

	router.GET("/health-check", handlers.HealthCheck)
	router.POST("/auth/sign-in", handler.SignIn)
	router.POST("/auth/sign-up", handler.SignUp)
	router.GET("/auth/me", middleware.RequireAuthMiddleware, handler.UserDetails)

	err = router.Run("localhost:8081")
	if err != nil {
		panic(err)
	}
}
