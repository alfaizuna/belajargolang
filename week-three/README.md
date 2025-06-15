# Belajar Go Lang - Week Three

## Deskripsi Proyek
Proyek ini adalah implementasi REST API sederhana menggunakan Go Lang. API ini menyediakan endpoint untuk mengelola task (CRUD operations).

## Struktur Proyek
```
task-api/
├── handler/     # Menangani HTTP request
├── model/       # Definisi struktur data
├── storage/     # Penyimpanan data (in-memory)
└── main.go      # Entry point aplikasi
```

## Teknologi yang Digunakan
- Go Lang
- Chi Router (untuk routing)
- JSON (untuk data format)

## Cara Menjalankan
1. Install Go Lang
2. Install dependencies:
   ```bash
   go get github.com/go-chi/chi/v5
   ```
3. Jalankan aplikasi:
   ```bash
   go run main.go
   ```
4. Server akan berjalan di `http://localhost:8080`

## Endpoint API
- `GET /tasks` - Mendapatkan semua task
- `POST /tasks` - Membuat task baru
- `GET /tasks/{id}` - Mendapatkan task berdasarkan ID
- `DELETE /tasks/{id}` - Menghapus task berdasarkan ID

## Format Data Task
```json
{
    "id": 1,
    "title": "Nama Task",
    "done": false
}
```

## Referensi Belajar Go Lang

### Dokumentasi Resmi
- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Go Tour](https://tour.golang.org/)

### Tutorial dan Kursus
- [Learn Go with Tests](https://github.com/quii/learn-go-with-tests)
- [Go Web Examples](https://gowebexamples.com/)
- [Go REST API Tutorial](https://tutorialedge.net/golang/creating-restful-api-with-golang/)

### Buku
- "Go Programming Language" oleh Alan A. A. Donovan & Brian W. Kernighan
- "Building Web Applications with Go" oleh Shiju Varghese

### Video Tutorial
- [Go Lang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU)
- [Go REST API Course](https://www.youtube.com/watch?v=SonwZ6MF5BE)

### Komunitas
- [Go Forum](https://forum.golangbridge.org/)
- [r/golang](https://www.reddit.com/r/golang/)
- [Gophers Slack](https://invite.slack.golangbridge.org/)

### Best Practices
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Proverbs](https://go-proverbs.github.io/)

## Konsep Go yang Dipelajari di Proyek Ini
1. Struct dan Interface
2. HTTP Server
3. JSON Handling
4. Routing
5. Error Handling
6. Package Management
7. Project Structure

## Tips Belajar Go
1. Mulai dengan dasar-dasar Go
2. Pelajari konsep concurrency
3. Praktikkan dengan membuat proyek kecil
4. Bergabung dengan komunitas Go
5. Baca kode sumber Go
6. Gunakan Go tools (go fmt, go vet, etc.)

## Kontribusi
Silakan buat pull request untuk kontribusi. Untuk perubahan besar, buka issue terlebih dahulu untuk mendiskusikan perubahan yang diinginkan.

## Lisensi
[MIT](https://choosealicense.com/licenses/mit/) 