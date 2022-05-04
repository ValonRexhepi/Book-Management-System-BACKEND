package tests

import (
	"testing"
	"time"

	"github.com/ValonRexhepi/Book-Management-System-BACKEND/controllers"
	"github.com/ValonRexhepi/Book-Management-System-BACKEND/models"
)

// TestUpdateBookSuccess test the successfull update of a book.
func TestUpdateBookSuccess(t *testing.T) {
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

	secondBook.TotalPages = 250
	err := controllers.UpdateBook(&secondBook)

	if err != nil {
		t.Errorf("Expected to update book but got %v", err)
	}

	updateBook, _ := controllers.GetBookByISBN("9780747538486")
	if secondBook.TotalPages != updateBook.TotalPages {
		t.Errorf("Expected to update total pages to 250 but got %v",
			updateBook.TotalPages)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}

// TestUpdateBookFail test the failed update of a book
func TestUpdateBookFail(t *testing.T) {
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

	secondBook.ISBN = firstBook.ISBN
	err := controllers.UpdateBook(&secondBook)

	if err == nil {
		t.Errorf("Expected to fail update book but got %v", err)
	}

	secondBook.ISBN = ""
	err = controllers.UpdateBook(&secondBook)

	if err == nil {
		t.Errorf("Expected to fail update book but got %v", err)
	}

	secondBook.Title = ""
	err = controllers.UpdateBook(&secondBook)

	if err == nil {
		t.Errorf("Expected to fail update book but got %v", err)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}
