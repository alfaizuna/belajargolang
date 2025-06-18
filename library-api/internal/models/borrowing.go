package models

import "time"

type Borrowing struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	BookID    uint       `json:"book_id" gorm:"not null"`
	Book      Book       `json:"book" gorm:"foreignKey:BookID"`
	StudentID uint       `json:"student_id" gorm:"not null"`
	Student   Student    `json:"student" gorm:"foreignKey:StudentID"`
	BorrowAt  time.Time  `json:"borrow_at" gorm:"not null"`
	ReturnAt  *time.Time `json:"return_at" gorm:"default:null"`
	Status    string     `json:"status" gorm:"type:varchar(20);not null"` // "borrowed", "returned"
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
