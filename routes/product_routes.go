package routes

import (
	"assignment23/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	r := router.Group("/products")
	{
		r.POST("/", controllers.CreateProduct)
		r.GET("/", controllers.GetProducts)
		r.GET("/:id", controllers.GetProductByID)
		r.PUT("/:id", controllers.UpdateProduct)
		r.DELETE("/:id", controllers.DeleteProduct)
	}
}
