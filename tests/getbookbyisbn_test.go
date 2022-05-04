package tests

import (
	"testing"
	"time"

	"github.com/ValonRexhepi/Book-Management-System-BACKEND/controllers"
	"github.com/ValonRexhepi/Book-Management-System-BACKEND/models"
)

// TestGetISBNBookSuccess test successfull retrieval of a book by isbn.
func TestGetISBNBookSuccess(t *testing.T) {
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

	bookByIsbn, err := controllers.GetBookByISBN("9780590353427")

	if err != nil {
		t.Errorf("Expected to success but got %v", err)
	}

	if !firstBook.Equal(bookByIsbn) {
		t.Errorf("Expected to have %v, but got %v", firstBook, bookByIsbn)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}

// TestGetISBNBookFail test the failed retrieval of a book by isbn.
func TestGetISBNBookFail(t *testing.T) {
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

	bookByIsbn, err := controllers.GetBookByISBN("15145214214")

	if err == nil || bookByIsbn == (&models.Book{}) {
		t.Errorf("Expected to fail but got %v", err)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}
