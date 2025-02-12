package main

import (
	"assignment23/config"
	"assignment23/models"
	"assignment23/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Memuat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Menghubungkan ke database
	config.ConnectDatabase()

	// Melakukan migrasi otomatis untuk memastikan tabel-tabel ada
	config.DB.AutoMigrate(&models.Product{}, &models.Inventory{}, &models.Order{}, &models.Image{})

	// Setup router untuk aplikasi
	r := routes.SetupRouter()

	// Menjalankan server di port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
