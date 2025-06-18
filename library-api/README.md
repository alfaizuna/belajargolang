# Library API

API sederhana untuk mengelola data buku dan mahasiswa menggunakan Go, PostgreSQL, dan Docker.

## Prasyarat

- Go 1.16 atau lebih baru
- Docker dan Docker Compose
- PostgreSQL (akan dijalankan melalui Docker)

## Cara Menjalankan

1. Clone repository ini
2. Jalankan PostgreSQL menggunakan Docker Compose:
   ```bash
   docker-compose up -d
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Jalankan aplikasi:
   ```bash
   go run cmd/main.go
   ```

## Endpoint API

### Books

- `POST /api/books` - Membuat buku baru
- `GET /api/books` - Mendapatkan semua buku
- `GET /api/books/:id` - Mendapatkan buku berdasarkan ID
- `PUT /api/books/:id` - Mengupdate buku
- `DELETE /api/books/:id` - Menghapus buku

### Students

- `POST /api/students` - Membuat data mahasiswa baru
- `GET /api/students` - Mendapatkan semua data mahasiswa
- `GET /api/students/:id` - Mendapatkan data mahasiswa berdasarkan ID
- `PUT /api/students/:id` - Mengupdate data mahasiswa
- `DELETE /api/students/:id` - Menghapus data mahasiswa

## Contoh Request

### Membuat Buku Baru
```bash
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan A. A. Donovan",
    "isbn": "978-0134190440",
    "published_at": "2024-03-20T00:00:00Z"
  }'
```

### Membuat Data Mahasiswa Baru
```bash
curl -X POST http://localhost:8080/api/students \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "nim": "2021001",
    "email": "john.doe@example.com",
    "major": "Computer Science"
  }'
```

# Get all books
curl -X GET http://localhost:8080/api/books

# Get book by ID
curl -X GET http://localhost:8080/api/books/1 