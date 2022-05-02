package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/ValonRexhepi/Book-Management-System-REST/controllers"
	"github.com/ValonRexhepi/Book-Management-System-REST/models"
)

// TestAddBook test successfully adding books in the database.
func TestAddBookSuccess(t *testing.T) {
	controllers.Connect()
	controllers.Migrate()

	var addBookSuccessTests = []models.Book{
		{
			Title:         "Harry Potter and the Philosopher's Stone",
			Author:        "J. K. Rowling",
			TotalPages:    320,
			PublishedDate: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
			ISBN:          "9780590353427",
		},
		{
			Title:         "Harry Potter and the Chamber of Secrets",
			Author:        "J. K. Rowling",
			TotalPages:    2511,
			PublishedDate: time.Date(1998, time.July, 2, 0, 0, 0, 0, time.UTC),
			ISBN:          "9780747538486",
		},
	}

	for _, tt := range addBookSuccessTests {
		testname := fmt.Sprintf("TEST:%s, %s", tt.Title, tt.ISBN)
		t.Run(testname, func(t *testing.T) {
			id, err := controllers.AddBook(&tt)
			if id == 0 || err != nil {
				t.Errorf("Expected success and got %v", err)
			}
		})
	}
	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}

// TestAddBook test failed adding books in the database.
func TestAddBookFail(t *testing.T) {
	controllers.Connect()
	controllers.Migrate()

	exampleBook := models.Book{
		Title:         "La Divina Comedia",
		Author:        "Dante Alighieri",
		TotalPages:    250,
		PublishedDate: time.Date(1300, time.July, 25, 0, 0, 0, 0, time.UTC),
		ISBN:          "9780747538486",
	}

	controllers.AddBook(&exampleBook)

	var addBookFailTests = []models.Book{
		{
			Title:         "La Divina Comedia",
			Author:        "Dante Alighieri",
			TotalPages:    250,
			PublishedDate: time.Date(1300, time.July, 25, 0, 0, 0, 0, time.UTC),
			ISBN:          "9780747538486",
		},
		{
			ID:            exampleBook.ID,
			Title:         "La Divina Comedia",
			Author:        "Dante Alighieri",
			TotalPages:    250,
			PublishedDate: time.Date(1300, time.July, 25, 0, 0, 0, 0, time.UTC),
			ISBN:          "978444988747538486",
		},
	}

	for _, tt := range addBookFailTests {
		testname := fmt.Sprintf("TEST:%s, %s", tt.Title, tt.ISBN)
		t.Run(testname, func(t *testing.T) {
			id, err := controllers.AddBook(&tt)
			if id != 0 || err == nil {
				t.Errorf("Expected fail and got %v", err)
			}
		})
	}
	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}
