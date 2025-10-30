package handlers

import (
	"cashierease/internal/models"
	"cashierease/internal/repositories"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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

func UploadGambarProduk(c *gin.Context) {
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

	file, err := c.FormFile("gambar_produk")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File 'gambar_produk' not found in request"})
		return
	}

	uploadDir := "./public/uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	extension := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("produk-%d-%d%s", id, time.Now().Unix(), extension)
	savePath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	urlPath := "/public/uploads/" + filename

	produk.GambarProduk = urlPath
	if err := repositories.UpdateProduk(&produk); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product path in database"})
		return
	}

	c.JSON(http.StatusOK, produk)
}