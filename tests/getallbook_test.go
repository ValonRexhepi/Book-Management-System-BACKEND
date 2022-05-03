package tests

import (
	"testing"
	"time"

	"github.com/ValonRexhepi/Book-Management-System-REST/controllers"
	"github.com/ValonRexhepi/Book-Management-System-REST/models"
)

// TestGetAllBook test successfully gets all books in the database.
func TestGetAllBookSuccess(t *testing.T) {
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

	secondBook := models.Book{
		Title:         "Harry Potter and the Chamber of Secrets",
		Author:        "J. K. Rowling",
		TotalPages:    2511,
		PublishedDate: time.Date(1998, time.July, 2, 0, 0, 0, 0, time.UTC),
		ISBN:          "9780747538486",
	}

	controllers.AddBook(&firstBook)
	controllers.AddBook(&secondBook)

	booksInDb, err := controllers.GetAllBooks()

	if err != nil {
		t.Errorf("Expected to success but got %v", err)
	}

	if !firstBook.Equal(booksInDb[0]) {
		t.Errorf("Expected to have %v, but got %v", firstBook, booksInDb[0])
	}

	if !secondBook.Equal(booksInDb[1]) {
		t.Errorf("Expected to have %v, but got %v", secondBook, booksInDb[1])
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}

// TestGetAllBook test failed geting all books in the database.
func TestGetAllBookFail(t *testing.T) {
	controllers.Connect()
	controllers.DB.Exec("DROP TABLE IF EXISTS books")
	controllers.Migrate()

	books, err := controllers.GetAllBooks()

	if err == nil && len(books) != 0 {
		t.Errorf("Expected to fail but got %v", err)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}
