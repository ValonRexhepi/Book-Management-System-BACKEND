package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ValonRexhepi/Book-Management-System-REST/controllers"
	"github.com/gin-gonic/gin"
)

// DeleteBookByID function to respond to a book delete in the database.
// Respond by JSON object with error and if error,
// else respond with success message.
func DeleteBookByID(c *gin.Context) {
	requestID, err := strconv.ParseUint(c.Param("bookid"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}

	err = controllers.DeleteBookByID(requestID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": fmt.Sprintf("Successfully deleted book with id : %v", requestID),
	})
}
