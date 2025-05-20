package handlers

import (
	"github.com/gin-gonic/gin"
	"goShortURL/internal/dto"
	"goShortURL/internal/models"
	"goShortURL/internal/repository"
	"goShortURL/internal/services"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Handler struct {
	authService services.AuthService
	userRepo    repository.UserRepository
}

func NewHandler(authService services.AuthService, userRepo repository.UserRepository) *Handler {
	return &Handler{
		authService: authService,
		userRepo:    userRepo,
	}
}

func (h *Handler) SignIn(c *gin.Context) {
	var userInput dto.UserInputSchema

	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := h.authService.SignIn(userInput.Email, userInput.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Login successful",
		"access_token": tokenString,
		"token_type":   "Bearer",
	})
}

func (h *Handler) SignUp(c *gin.Context) {
	var userInput dto.UserInputSchema

	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.userRepo.FindByEmail(userInput.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	user := &models.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: string(hashPassword),
	}

	if err := h.userRepo.Create(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *Handler) UserDetails(c *gin.Context) {
	userID, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := h.userRepo.FindByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
