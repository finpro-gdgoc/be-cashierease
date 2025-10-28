package handlers

import (
	"cashierease/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetToko(c *gin.Context) {
	toko, err := repositories.GetToko()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get toko info"})
		return
	}
	c.JSON(http.StatusOK, toko)
}

func UpdateToko(c *gin.Context) {
	toko, err := repositories.GetToko()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find toko info to update"})
		return
	}

	if err := c.ShouldBindJSON(&toko); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateToko(&toko); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update toko"})
		return
	}

	c.JSON(http.StatusOK, toko)
}