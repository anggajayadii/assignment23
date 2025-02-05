package controllers

import (
	"assignment23/config"
	"assignment23/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateOrder - Membuat pesanan baru
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah produk ada di stok
	var inventory models.Inventory
	if err := config.DB.Where("product_id = ?", order.ProductID).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak tersedia dalam stok"})
		return
	}

	// Cek apakah stok cukup
	if inventory.Quantity < order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stok tidak mencukupi"})
		return
	}

	// Kurangi stok
	inventory.Quantity -= order.Quantity
	config.DB.Save(&inventory)

	// Simpan pesanan ke database
	order.OrderDate = time.Now().Format("2006-01-02")
	config.DB.Create(&order)

	c.JSON(http.StatusCreated, gin.H{"message": "Pesanan berhasil dibuat", "order": order})
}

// GetOrderByID - Mendapatkan detail pesanan berdasarkan ID
func GetOrderByID(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID pesanan tidak valid"})
		return
	}

	var order models.Order
	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, order)
}
