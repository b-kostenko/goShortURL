package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", signIn)
		auth.POST("/sign-up", signUp)
	}
}
