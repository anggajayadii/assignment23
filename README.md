# assignment23
Web Server dan Golang Route

#Fitur
1. Manajemen Produk: Tambah, lihat, perbarui, dan hapus produk.
2. Manajemen Inventaris: Melihat dan memperbarui tingkat stok.
3. Manajemen Pesanan: Membuat dan mengambil detail pesanan.
4. RESTful API: Mendukung metode HTTP seperti GET, POST, PUT, dan DELETE menggunakan Gin.

#Teknologi yang Digunakan
Golang
Gin Framework
GORM (ORM untuk Golang)
MySQL
Postman (untuk pengujian API)


#Struktur Folder
│── config/                         # Konfigurasi database
    │── database.go                 
│── controllers/                    # Handler untuk setiap endpoint API
    │── inventory_controllers.go    
    │── order_controllers.go        
    │── product_controllers.go      
│── models/                         # Definisi model database
    │── models.go
│── routes/                         # Definisi routing
    │── routes.go
│── uploads/                        # Direktori penyimpanan gambar produk
│── main.go                         # Entry point aplikasi


#Endpoint API
1. Produk
GET /products → Lihat semua produk
GET /products/:id → Lihat detail produk berdasarkan ID
POST /products → Tambah produk baru
PUT /products/:id → Perbarui produk
DELETE /products/:id → Hapus produk

2. Inventaris
GET /inventory/:id → Lihat stok produk berdasarkan ID
PUT /inventory/:id → Perbarui stok produk

3. Pesanan
POST /orders → Buat pesanan baru
GET /orders/:id → Lihat detail pesanan

