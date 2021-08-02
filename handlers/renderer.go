package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func renderSuccess(c *gin.Context, code int) {
	c.JSON(code, gin.H{})
}

func renderFoundRecords(c *gin.Context, records interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": records,
	})
}

func renderError(c *gin.Context, err error) {
	var code int
	if err.Error() == "record not found" {
		code = http.StatusNotFound
	} else {
		code = http.StatusBadRequest
	}

	c.JSON(code, gin.H{
		"errors": []string{err.Error()},
	})
}

func renderErrors(c *gin.Context, errs []error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": func(errs []error) []string {
			strErrors := make([]string, len(errs))
			for i, err := range errs {
				strErrors[i] = err.Error()
			}
			return strErrors
		}(errs),
	})
}
