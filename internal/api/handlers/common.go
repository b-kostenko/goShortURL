package handlers

import (
	"github.com/gin-gonic/gin"
	"time"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "OK", "code": 200, "time": time.Now()})
}
