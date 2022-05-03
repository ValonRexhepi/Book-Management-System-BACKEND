package tests

import (
	"testing"
	"time"

	"github.com/ValonRexhepi/Book-Management-System-REST/controllers"
	"github.com/ValonRexhepi/Book-Management-System-REST/models"
)

// TestRemoveBook test successfully delete a book in the database.
func TestRemoveBookSucces(t *testing.T) {
	controllers.Connect()
	controllers.Migrate()

	firstBook := models.Book{
		Title:         "Harry Potter and the Philosopher's Stone",
		Author:        "J. K. Rowling",
		TotalPages:    320,
		PublishedDate: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
		ISBN:          "9780590353427",
	}

	controllers.AddBook(&firstBook)

	booksInDb, err := controllers.GetAllBooks()

	if len(booksInDb) != 1 || err != nil {
		t.Errorf("Expected to have one book but got %v with error: %v\n",
			len(booksInDb), err)
	}

	err = controllers.DeleteBook(firstBook.ID)

	if err != nil {
		t.Errorf("Expected to success delete but got %v", err)
	}

	booksInDb, err = controllers.GetAllBooks()
	if len(booksInDb) != 0 || err != nil {
		t.Errorf("Expected to have zero book but got %v with error: %v\n",
			len(booksInDb), err)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}

// TestRemoveBook test failed delete a book in the database.
func TestRemoveBookFail(t *testing.T) {
	controllers.Connect()
	controllers.Migrate()

	firstBook := models.Book{
		Title:         "Harry Potter and the Philosopher's Stone",
		Author:        "J. K. Rowling",
		TotalPages:    320,
		PublishedDate: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
		ISBN:          "9780590353427",
	}

	booksInDb, err := controllers.GetAllBooks()

	if len(booksInDb) != 0 || err != nil {
		t.Errorf("Expected to have zero book but got %v with error: %v\n",
			len(booksInDb), err)
	}

	err = controllers.DeleteBook(firstBook.ID)
	booksInDb, errGet := controllers.GetAllBooks()

	if (len(booksInDb) != 0 && err == nil) || errGet != nil {
		t.Errorf("Expected to fail delete but got %v", err)
	}

	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}
