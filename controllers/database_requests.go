package controllers

import (
	"github.com/ValonRexhepi/Book-Management-System-REST/models"
)

// AddBook function to add a new book to the database.
// Returns the id of the new book and nil if no error else returns the error.
func AddBook(bookToAdd *models.Book) (id uint, err error) {
	result := DB.Create(&bookToAdd)

	if result.Error != nil {
		return 0, result.Error
	}

	return bookToAdd.ID, nil
}

// UpdateBook function to update a book in the database.
// Returns nil if no error else returns the error.
func UpdateBook(bookToUpdate *models.Book) (err error) {
	result := DB.Save(&bookToUpdate)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllBooks function to return all book in the database.
// Returns an empty list and the error if error,
// else will return the list of books and nil.
func GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book
	result := DB.Find(&books)

	if result.Error != nil {
		return []*models.Book{}, result.Error
	}

	return books, nil
}
