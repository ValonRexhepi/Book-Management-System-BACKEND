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
	id, err := controllers.AddBook(&book)
	fmt.Println("Inserted ID", id)
	fmt.Println("Error", err)

	// Updating a book
	book.Author = "Valon Rexhepi"
	err = controllers.UpdateBook(&book)
	fmt.Println("Err updating", err)

	fmt.Println("App Started")
	for {

	}
}
