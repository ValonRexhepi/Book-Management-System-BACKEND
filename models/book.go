package models

import (
	"fmt"
	"time"
)

// Book struct that represents a Book.
type Book struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title         string    `json:"title" gorm:"not null"`
	Author        string    `json:"author" gorm:"not null"`
	TotalPages    uint      `json:"totalPages" gorm:"not null"`
	PublishedDate time.Time `json:"publishedDate" gorm:"not null"`
	ISBN          string    `json:"isbn" gorm:"unique;not null"`
}

// String method to return a string representation of the book.
func (b *Book) String() string {
	return fmt.Sprintf("{%v, %v, %v, %v, %v, %v}", b.ID, b.Title, b.Author,
		b.TotalPages, b.PublishedDate.Format("2006/01/02"), b.ISBN)
}

// Equal method to compare two Book objects return true if they are equal
// and false otherwise.
func (b *Book) Equal(other *Book) bool {
	if b.ID != other.ID || b.Title != other.Title || b.Author != other.Author ||
		b.TotalPages != other.TotalPages || b.ISBN != other.ISBN ||
		b.PublishedDate.UTC() != other.PublishedDate.UTC() {
		return false
	}

	return true
}
