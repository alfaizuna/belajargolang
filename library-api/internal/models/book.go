package models

import "time"

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Author      string    `json:"author" gorm:"not null"`
	ISBN        string    `json:"isbn" gorm:"unique;not null"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Book adalah model yang merepresentasikan data buku dalam sistem perpustakaan
// Field-field dalam struct Book:
// - ID: Primary key untuk identifikasi unik buku
// - Title: Judul buku (tidak boleh kosong)
// - Author: Penulis buku (tidak boleh kosong) 
// - ISBN: Nomor ISBN buku yang harus unik (tidak boleh kosong)
// - PublishedAt: Tanggal publikasi buku
// - CreatedAt: Timestamp kapan data buku dibuat
// - UpdatedAt: Timestamp kapan data buku terakhir diupdate

// Tag json digunakan untuk menentukan nama field saat data di-serialize ke JSON
// Tag gorm digunakan untuk konfigurasi database:
// - primaryKey: Menandakan field sebagai primary key
// - not null: Field tidak boleh kosong
// - unique: Nilai field harus unik

// Struct adalah tipe data komposit dalam Go yang memungkinkan kita untuk mengelompokkan 
// beberapa field dengan tipe data yang berbeda ke dalam satu unit. Struct mirip dengan 
// class di bahasa pemrograman lain, tapi tanpa inheritance. Struct digunakan untuk 
// membuat tipe data kustom yang dapat menyimpan data yang saling terkait.


