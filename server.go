package main

import (
	"ginDemo/controllers"
	"ginDemo/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.New()
	server.Use(middlewares.AccessLogger(), gin.Recovery())

	server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	videos := server.Group("/api/videos")
	{
		videos.GET("/", controllers.Find)
		videos.POST("/", controllers.Save)
		videos.PUT("/:id/", controllers.Update)
		videos.DELETE("/:id/", controllers.Delete)

	}
	server.Run(":8080")
}
