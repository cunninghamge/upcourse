package handlers

import (
	"fmt"
	"reflect"
)

func validateFields(s interface{}) []error {
	errs := []error{}

	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("validate")
		if tag == "onCreate" && v.Field(i).IsZero() {
			errs = append(errs, fmt.Errorf("%s is required", v.Type().Field(i).Name))
		}
	}

	return errs
}

func contains(slice []int, id int) bool {
	for _, num := range slice {
		if num == id {
			return true
		}
	}
	return false
}
