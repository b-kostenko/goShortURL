package shortlink

import (
	"github.com/gin-gonic/gin"
	"goShortURL/src/auth"
	"goShortURL/src/database"
	"math/rand"
	"net/http"
)

func listShortlinks(c *gin.Context) {
	//value, exist := c.Get("user")
	//
	//if !exist {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	//	return
	//}
	//user := value.(auth.User)

	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func createShortlink(c *gin.Context) {
	var userInput ShortlinkSchema
	var user auth.User
	value, exist := c.Get("user")
	if exist {
		user = value.(auth.User)
	}

	if err := c.BindJSON((&userInput)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	slug := createSlug(10)
	ShortURL := c.Request.Host + "/" + slug

	var shortLink = Shortlink{
		LongURL:  userInput.LongURL,
		Slug:     slug,
		ShortURL: ShortURL,
	}
	if exist {
		shortLink.UserID = &user.ID
	}
	result := database.DB.Create(&shortLink)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create shortlink"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shortlink created successfully", "details": ShortURL})
}

func createSlug(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
