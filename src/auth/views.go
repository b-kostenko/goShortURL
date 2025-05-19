package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"goShortURL/src/database"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

func signIn(c *gin.Context) {
	var user User
	var userInput UserInputSchema

	if err := c.BindJSON((&userInput)); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	result := database.DB.First(&user, "email = ?", userInput.Email)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email does not exist"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password is incorrect"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Login successful",
		"access_token": tokenString,
		"token_type":   "Bearer",
	})
}

func signUp(c *gin.Context) {
	var user User
	var userInput UserInputSchema

	if err := c.BindJSON((&userInput)); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error", err)
	}

	result := database.DB.First(&user, "email = ?", userInput.Email)
	if result.Error == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email is already registered"})
		return
	}

	user = User{Name: userInput.Name, Email: userInput.Email, Password: string(hashPassword)}
	result = database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func userDetails(c *gin.Context) {
	value, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user := value.(User)

	c.JSON(http.StatusOK, gin.H{
		"userID":    user.ID,
		"userName":  user.Name,
		"userEmail": user.Email,
	})
}
