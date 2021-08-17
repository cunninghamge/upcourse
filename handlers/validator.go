package handlers

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
)

func Validate(c *gin.Context, model interface{}, location string) []error {
	var errs []error

	if c.Request == nil {
		errs = append(errs, errors.New("invalid request"))
		return errs
	}

	jsonData, err := c.GetRawData()
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	jsonLoader := gojsonschema.NewBytesLoader(jsonData)
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + location)
	result, err := gojsonschema.Validate(schemaLoader, jsonLoader)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	if !result.Valid() {
		for _, e := range result.Errors() {
			errs = append(errs, errors.New(e.Description()))
		}
		return errs
	}

	if err := json.Unmarshal(jsonData, model); err != nil {
		errs = append(errs, err)
		return errs
	}

	return nil
}
