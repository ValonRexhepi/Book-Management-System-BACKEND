package controllers

import (
	"fmt"
	"strings"

	"github.com/ValonRexhepi/Book-Management-System-REST/models"
)

// DeleteBookByID function to delete a book in the database.
// Returns nil if no error else returns the error.
func DeleteBookByID(id uint64) (err error) {
	_, err = GetBookByID(id)

	if err != nil {
		return fmt.Errorf("no entry found in database for book with id: %v", id)
	}

	result := DB.Delete(&models.Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllBooks function to return all book in the database.
// Returns an empty list and the error if fail,
// else will return the list of books and nil.
func GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book
	result := DB.Find(&books)

	if result.Error != nil {
		return []*models.Book{}, result.Error
	}

	return books, nil
}

// GetBookByISBN function to return a book matching an ISBN in the database.
// Returns an empty book and the error if fail,
// else will return the book and nil.
func GetBookByISBN(isbn string) (*models.Book, error) {
	var book models.Book
	result := DB.Where("ISBN = ?", isbn).First(&book)

	if result.Error != nil {
		return &models.Book{}, result.Error
	}

	return &book, nil
}

// GetBookByID function to return a book matching an ID in the database.
// Returns an empty book and the error if fail,
// else will return the book and nil.
func GetBookByID(id uint64) (*models.Book, error) {
	var book models.Book
	result := DB.First(&book, id)

	if result.Error != nil {
		return &models.Book{}, result.Error
	}

	return &book, nil
}

// AddBook function to add a new book to the database.
// Returns the id of the new book and nil if no error else returns the error.
func AddBook(bookToAdd *models.Book) (id uint64, err error) {
	if len(strings.TrimSpace(bookToAdd.Title)) == 0 ||
		len(strings.TrimSpace(bookToAdd.Author)) == 0 ||
		len(strings.TrimSpace(bookToAdd.ISBN)) == 0 {
		return 0, fmt.Errorf("fill all fields")
	}

	result := DB.Omit("ID").Create(&bookToAdd)

	if result.Error != nil {
		return 0, result.Error
	}

	return bookToAdd.ID, nil
}

// UpdateBook function to update a book in the database.
// Returns nil if no error else returns the error.
func UpdateBook(bookToUpdate *models.Book) (err error) {
	if len(strings.TrimSpace(bookToUpdate.Title)) == 0 ||
		len(strings.TrimSpace(bookToUpdate.Author)) == 0 ||
		len(strings.TrimSpace(bookToUpdate.ISBN)) == 0 {
		return fmt.Errorf("fill all fields")
	}

	_, err = GetBookByID(bookToUpdate.ID)

	if err != nil {
		return fmt.Errorf("no entry found in database for book with id: %v", bookToUpdate.ID)
	}

	result := DB.Omit("ID").Save(&bookToUpdate)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
