package database

import (
	"fmt"
	"log"

	"library-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=library_db port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate
	err = db.AutoMigrate(&models.Book{}, &models.Student{}, &models.Borrowing{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
}

/*
File ini adalah package database untuk aplikasi Library API berbasis Go. 
Fungsinya sebagai pengelola koneksi dan inisialisasi database menggunakan GORM (ORM untuk Go) dan PostgreSQL.

Penjelasan bagian-bagian file:
- Import: Mengimpor package yang dibutuhkan, termasuk GORM, driver PostgreSQL, serta model-model data (Book, Student, Borrowing).
- Variabel global DB: Menyimpan instance *gorm.DB yang akan digunakan di seluruh aplikasi untuk operasi database.
- Fungsi InitDB():
    - Membuat koneksi ke database PostgreSQL dengan parameter yang sudah ditentukan (host, user, password, dbname, port, sslmode).
    - Jika koneksi gagal, aplikasi akan log.Fatal dan berhenti.
    - Melakukan auto-migrasi untuk model Book, Student, dan Borrowing agar tabel-tabel di database otomatis dibuat/sesuai dengan struktur model.
    - Menyimpan koneksi database ke variabel global DB.
    - Menampilkan pesan sukses jika koneksi berhasil.

File ini sangat penting karena menjadi fondasi seluruh operasi database pada aplikasi, memastikan koneksi dan struktur tabel sudah siap sebelum aplikasi berjalan.
*/

