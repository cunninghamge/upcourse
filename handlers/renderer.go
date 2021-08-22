package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
)

var errCodes = map[string]int{
	"record not found": http.StatusNotFound,
	"invalid request":  http.StatusBadRequest,
}

func renderErrors(c *gin.Context, errs ...error) {
	c.JSON(errorCode(errs), gin.H{
		"errors": func(errs []error) []string {
			strErrors := make([]string, len(errs))
			for i, err := range errs {
				strErrors[i] = err.Error()
			}
			return strErrors
		}(errs),
	})
}

func errorCode(errs []error) int {
	switch l := len(errs); {
	case l > 1 || strings.Contains(errs[0].Error(), "is required"):
		return http.StatusBadRequest
	case l == 1:
		if status, ok := errCodes[errs[0].Error()]; ok {
			return status
		}
	}
	return http.StatusInternalServerError
}

func renderFoundRecords(c *gin.Context, records interface{}) {
	payload, err := jsonapi.Marshal(records)
	if err != nil {
		renderErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, payload)
}
