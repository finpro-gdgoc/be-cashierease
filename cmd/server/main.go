package main

import (
	"cashierease/config"
	"cashierease/internal/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectDatabase()
}

func main() {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	routes.SetupAuthRoutes(apiV1)

	routes.SetupProdukRoutes(apiV1)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}