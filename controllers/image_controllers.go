package controllers

import (
	"assignment23/config"
	"assignment23/models"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	UploadDir   = "uploads/"
	MaxFileSize = 2 * 1024 * 1024 // 2MB
)

var allowedExtensions = map[string]bool{
	".png": true, ".jpg": true, ".jpeg": true,
}

// UploadImage handles image upload
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	if file.Size > MaxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the limit of 2MB"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format. Only PNG, JPG, and JPEG are allowed"})
		return
	}

	filePath := filepath.Join(UploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	productID := c.PostForm("product_id")
	image := models.Image{FileName: file.Filename}
	if productID != "" {
		id, err := strconv.ParseUint(productID, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product_id"})
			return
		}
		image.ProductID = uint(id)
	}

	config.DB.Create(&image)

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": file.Filename})
}

// GetImage serves image file based on Product ID
func GetImage(c *gin.Context) {
	var image models.Image
	productID := c.Param("product_id")

	if err := config.DB.Where("product_id = ?", productID).First(&image).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image not found for the given product ID"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	filePath := filepath.Join(UploadDir, image.FileName)
	c.File(filePath)
}
