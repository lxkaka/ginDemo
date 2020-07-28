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
	server.POST("/signup", controllers.SignUp)
	server.POST("/auth", controllers.Auth)

	videoController := controllers.VideoController{}

	videos := server.Group("/api/videos")
	videos.Use(middlewares.JWTAuth())
	{
		videos.GET("/", videoController.Find)
		videos.POST("/", videoController.Save)
		videos.PUT("/:id/", videoController.Update)
		videos.DELETE("/:id/", videoController.Delete)

	}
	server.Run(":8080")
}
