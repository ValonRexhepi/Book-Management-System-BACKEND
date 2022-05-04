package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/ValonRexhepi/Book-Management-System-BACKEND/controllers"
	"github.com/ValonRexhepi/Book-Management-System-BACKEND/models"
)

// TestAddBookSuccess test the successfull addition of new books.
func TestAddBookSuccess(t *testing.T) {
	controllers.Connect()
	controllers.DB.Exec("DROP TABLE IF EXISTS books")
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

// TestAddBookFail test the failed addition of new books.
func TestAddBookFail(t *testing.T) {
	controllers.Connect()
	controllers.DB.Exec("DROP TABLE IF EXISTS books")
	controllers.Migrate()

	var exampleBook = models.Book{

		Title:         "Harry Potter and the Philosopher's Stone",
		Author:        "J. K. Rowling",
		TotalPages:    320,
		PublishedDate: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
		ISBN:          "9780590353427",
	}

	controllers.AddBook(&exampleBook)

	var addBookSuccessFails = []models.Book{
		{

			Title:         "New Book With Same ISBN",
			Author:        "J. K. Rowling",
			TotalPages:    320,
			PublishedDate: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
			ISBN:          "9780590353427",
		},
		{
			Title:         "Harry Potter and the Philosopher's Stone",
			Author:        "",
			TotalPages:    320,
			PublishedDate: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
			ISBN:          "9780590353426",
		},
		{
			Title: "		   ",
			Author:        "J. K. Rowling",
			TotalPages:    2511,
			PublishedDate: time.Date(1998, time.July, 2, 0, 0, 0, 0, time.UTC),
			ISBN:          "    ",
		},
	}

	for _, tt := range addBookSuccessFails {
		testname := fmt.Sprintf("TEST:%d, %v", tt.TotalPages, tt.PublishedDate)
		t.Run(testname, func(t *testing.T) {
			id, err := controllers.AddBook(&tt)
			if id != 0 || err == nil {
				t.Errorf("Expected fail and got %v", err)
			}
		})
	}
	controllers.DB.Exec("DROP TABLE IF EXISTS books")
}
