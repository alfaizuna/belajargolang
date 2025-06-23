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

## 🔐 Authentication & Authorization (JWT, Middleware, RBAC)

### Fitur
- **Register & Login**: User dapat register dan login, mendapatkan JWT token.
- **JWT Middleware**: Setiap request ke endpoint yang butuh autentikasi harus menyertakan header `Authorization: Bearer <token>`.
- **RBAC (Role-Based Access Control)**: Endpoint tertentu hanya bisa diakses oleh role tertentu (misal: admin).
- **Profile**: Endpoint `/api/auth/profile` untuk melihat data user yang sedang login.

### Tools
- [`github.com/golang-jwt/jwt/v5`](https://github.com/golang-jwt/jwt) — Untuk generate dan validasi JWT.
- [`github.com/go-playground/validator`](https://github.com/go-playground/validator) — Untuk validasi input register/login.

### Cara Kerja
1. **Register**: POST ke `/api/auth/register` dengan body JSON username, email, password, role.
2. **Login**: POST ke `/api/auth/login` dengan email & password, response berisi JWT token.
3. **Akses endpoint protected**: Sertakan header `Authorization: Bearer <token>`.
4. **RBAC**: Endpoint tertentu (misal `/api/admin-only`) hanya bisa diakses oleh user dengan role `admin`.

### Contoh Request
```bash
# Register
curl -X POST http://localhost:8080/api/auth/register \
  -H 'Content-Type: application/json' \
  -d '{"username":"admin","email":"admin@mail.com","password":"admin123","role":"admin"}'

# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"admin@mail.com","password":"admin123"}'

# Profile (protected)
curl http://localhost:8080/api/auth/profile \
  -H 'Authorization: Bearer <token-dari-login>'
```

### Catatan
- JWT secret bisa diatur lewat environment variable `JWT_SECRET` (default: `secret`).
- Role yang didukung: `admin`, `student` (bisa dikembangkan sesuai kebutuhan).

## 🌱 Environment Variable

Aplikasi ini menggunakan environment variable untuk konfigurasi sensitif:
- `JWT_SECRET` — Secret key untuk JWT (default: `secret`)
- `DB_DSN` — Data Source Name untuk koneksi database PostgreSQL (default: lihat kode)

### Contoh .env
```
JWT_SECRET=secret
DB_DSN=host=localhost user=postgres password=postgres dbname=library_db port=5433 sslmode=disable
```

Gunakan package seperti [github.com/joho/godotenv](https://github.com/joho/godotenv) atau fitur bawaan deployment (Docker, CI/CD, dsb) untuk load file `.env` ke environment. 