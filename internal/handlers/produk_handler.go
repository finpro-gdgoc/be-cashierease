package handlers

import (
	"cashierease/internal/models"
	"cashierease/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func CreateProduk(c *gin.Context) {
	var produk models.Produk
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	produk.SlugProduk = slug.Make(produk.NamaProduk)

	if err := repositories.CreateProduk(&produk); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create produk"})
		return
	}

	c.JSON(http.StatusCreated, produk)
}

func GetAllProduk(c *gin.Context) {
	produks, err := repositories.GetAllProduk()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get produks"})
		return
	}
	c.JSON(http.StatusOK, produks)
}

func GetProdukById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	produk, err := repositories.GetProdukById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk not found"})
		return
	}
	c.JSON(http.StatusOK, produk)
}

func UpdateProduk(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	produk, err := repositories.GetProdukById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk not found"})
		return
	}

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	produk.SlugProduk = slug.Make(produk.NamaProduk)
    produk.ID = uint(id)

	if err := repositories.UpdateProduk(&produk); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update produk"})
		return
	}
	c.JSON(http.StatusOK, produk)
}

func DeleteProduk(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := repositories.DeleteProduk(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete produk"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produk deleted successfully"})
}