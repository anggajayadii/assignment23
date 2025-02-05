package controllers

import (
	"assignment23/config"
	"assignment23/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetInventory - Mendapatkan stok suatu produk berdasarkan ProductID
func GetInventory(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID produk tidak valid"})
		return
	}

	var inventory models.Inventory
	if err := config.DB.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stok tidak ditemukan untuk produk ini"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// UpdateInventory - Menambah atau mengurangi stok produk
func UpdateInventory(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID produk tidak valid"})
		return
	}

	var inventory models.Inventory
	if err := config.DB.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stok tidak ditemukan"})
		return
	}

	var input struct {
		Quantity int `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inventory.Quantity += input.Quantity
	config.DB.Save(&inventory)

	c.JSON(http.StatusOK, gin.H{"message": "Stok berhasil diperbarui", "inventory": inventory})
}
