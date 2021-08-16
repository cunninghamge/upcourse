package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
)

// TODO testing
func Validate(c *gin.Context, location string) ([]byte, []error) {
	var errs []error

	if c.Request == nil {
		errs = append(errs, errors.New("invalid request"))
		return nil, errs
	}

	jsonData, err := c.GetRawData()
	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}
	jsonLoader := gojsonschema.NewBytesLoader(jsonData)
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + location)

	result, err := gojsonschema.Validate(schemaLoader, jsonLoader)
	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}

	if !result.Valid() {
		for _, e := range result.Errors() {
			errs = append(errs, errors.New(e.Description()))
		}
		return nil, errs
	}

	return jsonData, nil
}
