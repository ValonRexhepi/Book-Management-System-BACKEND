package models

import "time"

// Book struct that represents a Book.
type Book struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title         string    `json:"title" gorm:"not null"`
	Author        string    `json:"author" gorm:"not null"`
	TotalPages    uint      `json:"totalPages" gorm:"not null"`
	PublishedDate time.Time `json:"publishedDate" gorm:"not null"`
	ISBN          string    `json:"isbn" gorm:"unique;not null"`
}
