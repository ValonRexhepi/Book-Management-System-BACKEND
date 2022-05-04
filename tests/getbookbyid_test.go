package tests

import (
	"testing"
	"time"

	"github.com/ValonRexhepi/Book-Management-System-REST/controllers"
	"github.com/ValonRexhepi/Book-Management-System-REST/models"
)

// TestGetIDBookSuccess test the successfull retrieval of a book by id.
func TestGetIDBookSuccess(t *testing.T) {
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

	bookByIsbn, err := controllers.GetBookByID(1)

	if err != nil {
		t.Errorf("Expected to success but got %v", err)
	}

	if !firstBook.Equal(bookByIsbn) {
		t.Errorf("Expected to have %v, but got %v", firstBook, bookByIsbn)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}

// TestGetIDBookFail test the failed retrieval of a book by id..
func TestGetIDBookFail(t *testing.T) {
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

	controllers.AddBook(&firstBook) // id will be 1

	bookByIsbn, err := controllers.GetBookByID(2)

	if err == nil || bookByIsbn == (&models.Book{}) {
		t.Errorf("Expected to fail but got %v", err)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}
