package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
)

const (
	ErrNotFound   = "record not found"
	ErrBadRequest = "invalid request"
)

func renderFoundRecords(c *gin.Context, records interface{}) {
	payload, err := jsonapi.Marshal(records)
	if err != nil {
		renderError(c, err)
		return
	}

	c.JSON(http.StatusOK, payload)
}

func renderError(c *gin.Context, err error) {
	var code int
	if err.Error() == ErrNotFound {
		code = http.StatusNotFound
	} else if err.Error() == ErrBadRequest {
		code = http.StatusBadRequest
	} else {
		code = http.StatusInternalServerError
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
