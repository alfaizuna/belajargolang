package main

/*
INTI PELAJARAN:
1. REST API dengan Go
   - Membuat API sederhana dengan 4 endpoint (GET, POST, GET by ID, DELETE)
   - Menggunakan Chi Router: ringan, cepat, dan sintaks yang bersih
   - Implementasi CRUD (Create, Read, Update, Delete) dasar

2. Struktur Proyek
   - Pemisahan kode menjadi package (handler, model, storage)
   - Penggunaan modular untuk maintainability yang lebih baik
   - Organisasi kode yang rapi dan terstruktur

3. HTTP Server
   - Cara membuat server HTTP di Go
   - Penanganan routing dengan method yang berbeda
   - Penggunaan parameter dinamis dalam URL

4. Best Practices
   - Penggunaan Chi Router: performa tinggi, middleware fleksibel, sub-routing
   - Pemisahan concerns (handler, model, storage)
   - Penamaan endpoint yang RESTful
*/

import (
	"fmt"
	"net/http"
	"task-api/handler"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/tasks", handler.GetTasks)
	r.HandleFunc("/tasks", handler.CreateTask)
	r.HandleFunc("/tasks/{id}", handler.GetTaskByID)
	r.HandleFunc("/tasks/{id}", handler.DeleteTask)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
