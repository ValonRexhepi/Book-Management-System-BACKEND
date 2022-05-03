package tests

import (
	"testing"
	"time"

	"github.com/ValonRexhepi/Book-Management-System-REST/controllers"
	"github.com/ValonRexhepi/Book-Management-System-REST/models"
)

// TestGetIDBook test successfully gets a book by ID in the database.
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

// TestGetIDBook test failed geting a book by ID in the database.
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

	controllers.AddBook(&firstBook)

	bookByIsbn, err := controllers.GetBookByID(2)

	if err == nil || bookByIsbn == (&models.Book{}) {
		t.Errorf("Expected to fail but got %v", err)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}
