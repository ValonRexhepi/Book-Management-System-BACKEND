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

	books, _ := controllers.GetAllBooks()

	// for _, book := range books {
	// 	fmt.Printf("Book finded : %v\n", book)
	// }

	fmt.Printf("Book finded : %+v\n", books)

	fmt.Println("App Started")
	controllers.DB.Exec("DROP TABLE IF EXISTS books")

	for {

	}
}
