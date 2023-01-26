package main

import (
	routes "auth-jwt-gin/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port = os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H("success": "Access berhasil api 1"))
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H("success": "Access berhasil api 2"))
	})

	router.RUN(":" + port)
}
