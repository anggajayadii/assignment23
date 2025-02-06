package controllers

import (
	"assignment23/config"
	"assignment23/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Order (Menambahkan pesanan baru)
func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah produk ada
	var product models.Product
	if err := config.DB.First(&product, order.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	// Simpan order ke database
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pesanan berhasil dibuat", "order": order})
}

// Get All Orders (Menampilkan semua pesanan)
func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Preload("Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// Get Order by ID (Menampilkan pesanan berdasarkan ID)
func GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")
	var order models.Order

	if err := config.DB.Preload("Product").First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// Update Order (Memperbarui pesanan)
func UpdateOrder(c *gin.Context) {
	orderID := c.Param("id")
	var order models.Order

	// Cek apakah pesanan ada
	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
		return
	}

	// Bind data JSON ke struct order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan perubahan
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pesanan berhasil diperbarui", "order": order})
}

// Delete Order (Menghapus pesanan)
func DeleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	var order models.Order

	// Cek apakah pesanan ada
	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
		return
	}

	// Hapus pesanan
	if err := config.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pesanan berhasil dihapus"})
}
