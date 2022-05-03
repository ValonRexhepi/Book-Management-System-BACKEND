package main

import (
	"fmt"
	"time"

	"github.com/ValonRexhepi/Book-Management-System-REST/controllers"
	"github.com/ValonRexhepi/Book-Management-System-REST/models"
)

func main() {
	controllers.Connect()
	controllers.Migrate()

	// Adding a book
	book := models.Book{
		Title:         "Test title",
		Author:        "Test Author",
		TotalPages:    100,
		PublishedDate: time.Now(),
		ISBN:          "asads",
	}
	book2 := models.Book{
		Title:         "Test title 2",
		Author:        "Test Author 2",
		TotalPages:    100,
		PublishedDate: time.Now(),
		ISBN:          "asads2",
	}
	controllers.AddBook(&book)
	controllers.AddBook(&book2)

	// Updating a book
	book.Author = "Valon Rexhepi"
	controllers.UpdateBook(&book)

	controllers.DeleteBook(&book)

	// Get all books
	books, _ := controllers.GetAllBooks()
	fmt.Printf("Books finded : %+v\n", books)

	// Get book by ISBN
	bookByIsbn, _ := controllers.GetBookByISBN("asads2")
	fmt.Printf("Book by ISBN finded : %+v\n", bookByIsbn)

	// Get book by ISBN
	bookByID, _ := controllers.GetBookByID(2)
	fmt.Printf("Book by ID finded : %+v\n", bookByID)

	fmt.Println("App Started")

	controllers.DB.Exec("DROP TABLE IF EXISTS books")

}
