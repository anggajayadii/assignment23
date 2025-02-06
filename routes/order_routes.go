package routes

import (
	"assignment23/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	r := router.Group("/orders")
	{
		r.POST("/", controllers.CreateOrder)      // Menambahkan pesanan
		r.GET("/", controllers.GetAllOrders)      // Mendapatkan semua pesanan
		r.GET("/:id", controllers.GetOrderByID)   // Mendapatkan pesanan berdasarkan ID
		r.PUT("/:id", controllers.UpdateOrder)    // Memperbarui pesanan
		r.DELETE("/:id", controllers.DeleteOrder) // Menghapus pesanan
	}
}
