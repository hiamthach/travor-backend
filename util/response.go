package util

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

// I want to create a util response function that will return a JSON response with a status code and a message, data.
func SuccessResponse(message string, data interface{}) gin.H {
	return gin.H{
		"message": message,
		"data":    data,
	}
}
