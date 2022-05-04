package server

import (
	"fmt"
	"os"

	"github.com/ValonRexhepi/Book-Management-System-BACKEND/routes"
	"github.com/gin-gonic/gin"
)

// LaunchServer function to define default router, define the routes
// and run the server on port localhost:8080.
func LaunchServer() {
	router := gin.Default()

	router.DELETE("/books/deletebook/:bookid", routes.DeleteBookByID)

	router.GET("/books", routes.GetAllBooks)
	router.GET("/books/getbookbyid/:bookid", routes.GetBookByID)
	router.GET("/books/getbookbyisbn/:bookisbn", routes.GetBookByISBN)

	router.POST("/books/updatebook", routes.UpdateBook)
	router.POST("/books", routes.AddBook)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Can't launch the server: ", err)
		os.Exit(-1)
	}
}
