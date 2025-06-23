package main

import (
	"log"
	"net/http"

	"library-api/internal/handlers"
	"library-api/internal/middleware"
	"library-api/internal/models"
	"library-api/internal/repository"
	"library-api/pkg/database"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, menggunakan environment variable dari sistem")
	}

	// Initialize database
	database.InitDB()

	// Create router
	r := chi.NewRouter()

	// Middleware
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.RealIP)

	// Initialize handlers
	bookHandler := handlers.NewBookHandler()
	studentHandler := handlers.NewStudentHandler()
	borrowingHandler := handlers.NewBorrowingHandler()

	// Tambah: UserRepository dan AuthHandler
	userRepo := repository.NewUserRepository(database.DB)
	authHandler := handlers.NewAuthHandler(userRepo)

	// Book routes
	r.Route("/api/books", func(r chi.Router) {
		r.Post("/", bookHandler.CreateBook)
		r.Get("/", bookHandler.GetBooks)
		r.Get("/{id}", bookHandler.GetBook)
		r.Put("/{id}", bookHandler.UpdateBook)
		r.Delete("/{id}", bookHandler.DeleteBook)
	})

	// Student routes
	r.Route("/api/students", func(r chi.Router) {
		r.Post("/", studentHandler.CreateStudent)
		r.Get("/", studentHandler.GetStudents)
		r.Get("/{id}", studentHandler.GetStudent)
		r.Put("/{id}", studentHandler.UpdateStudent)
		r.Delete("/{id}", studentHandler.DeleteStudent)
	})

	// Borrowing routes
	r.Route("/api/borrowings", func(r chi.Router) {
		r.Post("/", borrowingHandler.BorrowBook)
		r.Get("/", borrowingHandler.GetBorrowings)
		r.Put("/{id}/return", borrowingHandler.ReturnBook)
		r.Get("/student/{student_id}", borrowingHandler.GetStudentBorrowings)
		r.Get("/book/{book_id}", borrowingHandler.GetBookBorrowings)
	})

	// Auth routes
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(userRepo))
			r.Get("/profile", authHandler.Profile)
		})
	})

	// Contoh RBAC: hanya admin bisa akses endpoint ini
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(userRepo))
		r.Use(middleware.RequireRole(models.AdminRole))
		r.Get("/api/admin-only", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, admin!"))
		})
	})

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
