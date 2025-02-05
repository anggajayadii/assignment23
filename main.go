package main

import (
	"assignment23/config"
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

	r := gin.Default()

	config.ConnectDatabase()

	routes.ProductRoutes(r)

	r.Run(":8080")
}
