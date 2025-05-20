package router

import "github.com/gin-gonic/gin"

func AuthRouters(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", signIn)
		auth.POST("/sign-up", signUp)
		auth.GET("/me", RequireAuthMiddleware, userDetails)
	}
}
