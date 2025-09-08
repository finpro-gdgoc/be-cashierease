package main

import (
	"cashierease/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}