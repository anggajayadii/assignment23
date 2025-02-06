package controllers

import (
	"assignment23/config"
	"assignment23/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 📌 1️⃣ Menambahkan stok baru
func AddInventory(c *gin.Context) {
	var inventory models.Inventory

	// Bind JSON ke struct Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah produk dengan ProductID ada di tabel products
	var product models.Product
	if err := config.DB.First(&product, inventory.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	// Simpan data ke inventory
	config.DB.Create(&inventory)

	// Response sukses
	c.JSON(http.StatusCreated, gin.H{
		"message":   "Stok berhasil ditambahkan",
		"inventory": inventory,
	})
}

// 📌 2️⃣ Melihat stok berdasarkan `product_id`
func GetInventoryByProductID(c *gin.Context) {
	productID := c.Param("product_id")

	var inventory []models.Inventory
	if err := config.DB.Preload("Product").Where("product_id = ?", productID).Find(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(inventory) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Stok tidak ditemukan untuk produk ini"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"inventory": inventory,
	})
}

// 📌 3️⃣ Memperbarui stok (Menambah/Mengurangi quantity)
func UpdateInventory(c *gin.Context) {
	inventoryID := c.Param("id") // Ambil inventory ID dari parameter URL

	var inventory models.Inventory
	if err := config.DB.First(&inventory, inventoryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stok tidak ditemukan"})
		return
	}

	// Bind data dari request ke struct Inventory
	var input struct {
		Quantity int    `json:"quantity"`
		Location string `json:"location"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update stok (quantity & location jika ada perubahan)
	inventory.Quantity = input.Quantity
	inventory.Location = input.Location
	config.DB.Save(&inventory)

	// Response sukses
	c.JSON(http.StatusOK, gin.H{
		"message":   "Stok berhasil diperbarui",
		"inventory": inventory,
	})
}

// 📌 4️⃣ Menghapus stok berdasarkan `id`
func DeleteInventory(c *gin.Context) {
	inventoryID := c.Param("id") // Ambil inventory ID dari parameter URL

	var inventory models.Inventory
	if err := config.DB.First(&inventory, inventoryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stok tidak ditemukan"})
		return
	}

	// Hapus stok
	config.DB.Delete(&inventory)

	// Response sukses
	c.JSON(http.StatusOK, gin.H{"message": "Stok berhasil dihapus"})
}
