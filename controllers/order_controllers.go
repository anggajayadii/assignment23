package controllers

import (
	"assignment23/config"
	"assignment23/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var orderInput struct {
		ProductID int    `json:"product_id"`
		Quantity  int    `json:"quantity"`
		OrderDate string `json:"order_date"`
	}

	// Bind JSON untuk mengambil input pesanan
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi produk
	var product models.Product
	if err := config.DB.First(&product, orderInput.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Validasi jumlah stok
	var inventory models.Inventory
	if err := config.DB.Where("product_id = ?", orderInput.ProductID).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	// Pastikan stok mencukupi untuk pesanan
	if inventory.Quantity < orderInput.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stok tidak cukup"})
		return
	}

	// Mengurangi stok setelah pemesanan
	inventory.Quantity -= orderInput.Quantity
	config.DB.Save(&inventory)

	// Membuat entri pesanan
	order := models.Order{
		ProductID: orderInput.ProductID,
		Quantity:  orderInput.Quantity,
		OrderDate: orderInput.OrderDate,
	}

	// Simpan pesanan ke database
	config.DB.Create(&order)

	c.JSON(http.StatusOK, order)
}

// GetOrderByID - Mengambil detail pesanan berdasarkan ID

func GetOrderByID(c *gin.Context) {
	// Konversi orderID dari string ke uint
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order

	// Query dengan Preload hanya jika relasi benar-benar ada
	if err := config.DB.Preload("Product").Where("id = ?", orderID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
