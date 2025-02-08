package controllers

import (
	"assignment23/config"
	"assignment23/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 📌 1️⃣ Menambahkan stok baru
func GetInventory(c *gin.Context) {
	productID := c.Param("product_id")
	var inventory models.Inventory

	// Gunakan Preload untuk mengambil data Product juga
	if err := config.DB.Preload("Product").Find(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventory"})
		return
	}

	// Cari inventory berdasarkan product_id
	if err := config.DB.Where("product_id = ?", productID).Find(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// UpdateInventory - Memperbarui stok produk (menambah atau mengurangi)
func UpdateInventory(c *gin.Context) {
	var inventory models.Inventory
	productID := c.Param("product_id")

	// Cek apakah produk ada di database inventory
	if err := config.DB.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found in inventory"})
		return
	}

	// Data dari request body (JSON)
	var input struct {
		Quantity int    `json:"quantity"`
		Location string `json:"location"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update stok dan lokasi
	inventory.Quantity = input.Quantity
	inventory.Location = input.Location

	// Simpan perubahan ke database
	if err := config.DB.Save(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Inventory updated successfully",
		"inventory": inventory,
	})
}
