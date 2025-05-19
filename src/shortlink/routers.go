package shortlink

import (
	"github.com/gin-gonic/gin"
	"goShortURL/src/auth"
)

func Routers(router *gin.Engine) {
	shortlinks := router.Group("/shortlinks")
	shortlinks.Use(auth.RequireAuthMiddleware)
	{
		shortlinks.GET("/", listShortlinks)
		//shortlinks.GET("/:id", getShortlink)
		shortlinks.POST("/", createShortlink)
		//shortlinks.PUT("/:id", updateShortlink)
		//shortlinks.PATCH("/:id", patchShortlink)
		//shortlinks.DELETE("/:id", deleteShortlink)
	}
}
