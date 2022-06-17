// a package for validating things (like structs, json, etc) based on github.com/go-playground/validator/v10
package validator

import (
	"encoding/json"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var (
	// validate is a single instance of Validator
	validate *validator.Validate
)

func init() {
	validate = validator.New() // initialize a single instance of Validator
}

// Struct validates a struct based on the tags
func Struct(s any) error {
	return validate.Struct(s)
}

// Var validates a variable based on given validator tags
func Var(v any, validators string) error {
	return validate.Var(v, validators)
}

// unmarshal a json and validate it based on the tags
func JSON(j []byte, t reflect.Type) (error, any) {
	v := reflect.Zero(t).Interface()
	err := json.Unmarshal(j, &v)
	if err != nil {
		return err, nil
	}
	err = validate.Struct(v)
	return err, v
}
