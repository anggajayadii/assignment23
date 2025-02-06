package main

import (
	"assignment23/config"
	"assignment23/models"
	"assignment23/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Product{}, &models.Inventory{}, &models.Order{})

	r := gin.Default()

	config.ConnectDatabase()

	routes.ProductRoutes(r)
	routes.InventoryRoutes(r)
	routes.OrderRoutes(r)

	r.Run(":8080")
}
