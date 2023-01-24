package controllers

import (
	"github.com/gin-gonic/gin"
)

type EventController struct{}

func RegisterRoutes(router *gin.Engine) {

	page := router.Group("/pages")
	{
		page.GET("/:title", GetPageByTitle)
		page.GET("/", GetPage)
		page.POST("/", PostPage)
		page.DELETE("/:title", DeletePage)
		// page.PATCH("/:title", PatchPage)
		// page.PUT("/:title", PutPage)
	}

	// Add a healh check endpoint
	router.GET("/health", HealthCheck)
}
