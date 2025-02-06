package routes

import (
	"assignment23/controllers"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(router *gin.Engine) {
	r := router.Group("/inventory")
	{
		r.POST("/", controllers.AddInventory)                      // Tambah stok
		r.GET("/:product_id", controllers.GetInventoryByProductID) // Lihat stok berdasarkan product_id
		r.PUT("/:id", controllers.UpdateInventory)                 // Perbarui stok
		r.DELETE("/:id", controllers.DeleteInventory)              // Hapus stok
	}
}
