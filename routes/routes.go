package routes

import (
	"assignment23/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Produk Routes
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductByID)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	// Inventory Routes
	r.GET("/inventory/:product_id", controllers.GetInventory)
	r.PUT("/inventory/:product_id", controllers.UpdateInventory)

	// Orders Routes
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders/:id", controllers.GetOrderByID)

	// Image Routes
	r.POST("/upload", controllers.UploadImage)
	r.GET("/images/:product_id", controllers.GetImage)

	return r
}
