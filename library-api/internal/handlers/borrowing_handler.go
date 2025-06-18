package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"library-api/internal/models"
	"library-api/pkg/database"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type BorrowingHandler struct {
	db *gorm.DB
}

func NewBorrowingHandler() *BorrowingHandler {
	return &BorrowingHandler{
		db: database.DB,
	}
}

// BorrowBook handles the borrowing of a book
func (h *BorrowingHandler) BorrowBook(w http.ResponseWriter, r *http.Request) {
	var borrowing models.Borrowing
	if err := json.NewDecoder(r.Body).Decode(&borrowing); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if book exists
	var book models.Book
	if err := h.db.First(&book, borrowing.BookID).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Check if student exists
	var student models.Student
	if err := h.db.First(&student, borrowing.StudentID).Error; err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	// Check if book is already borrowed
	var existingBorrowing models.Borrowing
	if err := h.db.Where("book_id = ? AND status = ?", borrowing.BookID, "borrowed").First(&existingBorrowing).Error; err == nil {
		http.Error(w, "Book is already borrowed", http.StatusBadRequest)
		return
	}

	// Set borrow time and status
	borrowing.BorrowAt = time.Now()
	borrowing.Status = "borrowed"
	borrowing.ReturnAt = nil

	if err := h.db.Create(&borrowing).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Load book and student data
	h.db.Preload("Book").Preload("Student").First(&borrowing, borrowing.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(borrowing)
}

// ReturnBook handles the return of a book
func (h *BorrowingHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var borrowing models.Borrowing

	if err := h.db.First(&borrowing, id).Error; err != nil {
		http.Error(w, "Borrowing record not found", http.StatusNotFound)
		return
	}

	if borrowing.Status == "returned" {
		http.Error(w, "Book is already returned", http.StatusBadRequest)
		return
	}

	// Set return time
	now := time.Now()
	borrowing.ReturnAt = &now
	borrowing.Status = "returned"

	if err := h.db.Save(&borrowing).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Load book and student data
	h.db.Preload("Book").Preload("Student").First(&borrowing, borrowing.ID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(borrowing)
}

// GetBorrowings returns all borrowing records
func (h *BorrowingHandler) GetBorrowings(w http.ResponseWriter, r *http.Request) {
	var borrowings []models.Borrowing
	if err := h.db.Preload("Book").Preload("Student").Find(&borrowings).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(borrowings)
}

// GetStudentBorrowings returns all borrowing records for a specific student
func (h *BorrowingHandler) GetStudentBorrowings(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "student_id")
	var borrowings []models.Borrowing

	if err := h.db.Preload("Book").Preload("Student").Where("student_id = ?", studentID).Find(&borrowings).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(borrowings)
}

// GetBookBorrowings returns all borrowing records for a specific book
func (h *BorrowingHandler) GetBookBorrowings(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "book_id")
	var borrowings []models.Borrowing

	if err := h.db.Preload("Book").Preload("Student").Where("book_id = ?", bookID).Find(&borrowings).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(borrowings)
}
