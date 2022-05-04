package tests

import (
	"testing"
	"time"

	"github.com/ValonRexhepi/Book-Management-System-BACKEND/controllers"
	"github.com/ValonRexhepi/Book-Management-System-BACKEND/models"
)

// TestRemoveBookSuccess test the successfull deletion of a book.
func TestRemoveBookSuccess(t *testing.T) {
	controllers.Connect()
	controllers.DB.Exec("DROP TABLE IF EXISTS books")
	controllers.Migrate()

	firstBook := models.Book{
		Title:         "Harry Potter and the Philosopher's Stone",
		Author:        "J. K. Rowling",
		TotalPages:    320,
		PublishedDate: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
		ISBN:          "9780590353427",
	}

	controllers.AddBook(&firstBook)

	err := controllers.DeleteBookByID(firstBook.ID)

	if err != nil {
		t.Errorf("Expected to success delete but got %v", err)
	}

	booksInDb, err := controllers.GetAllBooks()
	if len(booksInDb) != 0 || err != nil {
		t.Errorf("Expected to have zero book but got %v books with error: %v\n",
			len(booksInDb), err)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}

// TestRemoveBookFail test the failed deletion of a book.
func TestRemoveBookFail(t *testing.T) {
	controllers.Connect()
	controllers.DB.Exec("DROP TABLE IF EXISTS books")
	controllers.Migrate()

	err := controllers.DeleteBookByID(1)
	booksInDb, errGet := controllers.GetAllBooks()

	if (len(booksInDb) != 0 && err == nil) || errGet != nil {
		t.Errorf("Expected to fail delete but got %v", err)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}
