package routes

import (
	"fmt"
	"net/http"

	"github.com/ValonRexhepi/Book-Management-System-REST/controllers"
	"github.com/ValonRexhepi/Book-Management-System-REST/models"
	"github.com/gin-gonic/gin"
)

// AddBook function to respond to a book added to the database.
// Respond by JSON object with error if error,
// else respond with success message and new book.
func AddBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}

	id, err := controllers.AddBook(&newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   fmt.Sprintf("Successfully added book with id : %v", id),
		"bookAdded": newBook,
	})

}

// UpdateBook function to respond to a book updated in the database.
// Respond by JSON object with error if error,
// else respond with success message and updated book.
func UpdateBook(c *gin.Context) {
	var bookToUpdate models.Book

	if err := c.BindJSON(&bookToUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}

	err := controllers.UpdateBook(&bookToUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     "Successfully updated book",
		"bookUpdated": bookToUpdate,
	})
}
