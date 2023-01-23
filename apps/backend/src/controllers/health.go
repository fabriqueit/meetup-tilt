package controllers

import "github.com/gin-gonic/gin"

// HealthCheck returns a simple 200 with message OK.
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
