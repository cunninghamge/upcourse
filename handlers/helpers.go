package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func renderSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": message,
	})
}

func renderFound(c *gin.Context, records interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": message,
		"data":    records,
	})
}

func renderNotFound(c *gin.Context, err error) {
	log.Printf("Error retrieving record from database.\nReason: %v", err)
	c.JSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
		"errors": "Record not found",
	})
}

func renderCreated(c *gin.Context, message string) {
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": message,
	})
}

func renderBindError(c *gin.Context, bindErr error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"error":  bindErr.Error(),
	})
}

func renderError(c *gin.Context, err error) {
	c.JSON(http.StatusServiceUnavailable, gin.H{
		"status": http.StatusServiceUnavailable,
		"error":  err,
	})
}

func renderErrors(c *gin.Context, errs []error) {
	c.JSON(http.StatusServiceUnavailable, gin.H{
		"status": http.StatusServiceUnavailable,
		"errors": func(errs []error) []string {
			strErrors := make([]string, len(errs))
			for i, err := range errs {
				strErrors[i] = err.Error()
			}
			return strErrors
		}(errs),
	})
}
