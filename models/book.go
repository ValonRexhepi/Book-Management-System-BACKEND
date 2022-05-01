package models

import "time"

// Book struct that represents a Book.
type Book struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	TotalPages    uint      `json:"totalPages"`
	PublishedDate time.Time `json:"publishedDate"`
	ISBN          string    `json:"isbn" gorm:"unique"`
}
