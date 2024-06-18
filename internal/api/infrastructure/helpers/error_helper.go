package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// HandleError handles errors by logging the error message and returning a JSON response
// with the error message and a specified status code.
//
// Parameters:
// - c: a pointer to a gin.Context object representing the current HTTP request context.
// - err: an error object representing the error to be handled.
// - statusCode: an integer representing the HTTP status code to be returned.
//
// Returns: None.
func HandleError(c *gin.Context, err error, statusCode int) {
	log.Printf("Error: %s", err.Error())

	data := map[string]interface{}{
		"error": err.Error(),
	}

	c.JSON(statusCode, data)
}
