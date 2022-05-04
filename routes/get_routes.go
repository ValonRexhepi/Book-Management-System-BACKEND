package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ValonRexhepi/Book-Management-System-REST/controllers"
	"github.com/gin-gonic/gin"
)

// GetAllBooks function to respond to a get request for all books.
// Respond by JSON object with error and empty list if error,
// else respond with the list of books.
func GetAllBooks(c *gin.Context) {
	books, err := controllers.GetAllBooks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
			"books": books,
		})
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetBookByISBN function to respond to a get request for a book
// matching an ISBN. Respond by JSON object with error and empty book if error,
// else respond with the book.
func GetBookByISBN(c *gin.Context) {
	requestISBN := c.Param("bookisbn")
	book, err := controllers.GetBookByISBN(requestISBN)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
			"book":  book,
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

// GetBookByID function to respond to a get request for a book matching an ID.
// Respond by JSON object with error and empty book if error,
// else respond with the book.
func GetBookByID(c *gin.Context) {
	requestID, err := strconv.ParseUint(c.Param("bookid"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}

	book, err := controllers.GetBookByID(requestID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
			"book":  book,
		})
		return
	}
	c.JSON(http.StatusOK, book)
}
