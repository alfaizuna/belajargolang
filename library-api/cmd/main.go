package main

import (
	"log"
	"net/http"

	"library-api/internal/handlers"
	"library-api/pkg/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize database
	database.InitDB()

	// Create router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	// Initialize handlers
	bookHandler := handlers.NewBookHandler()
	studentHandler := handlers.NewStudentHandler()
	borrowingHandler := handlers.NewBorrowingHandler()

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

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
