// a package for validating things (like structs, json, etc) based on github.com/go-playground/validator/v10
package validator

import (
	"encoding/json"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func Struct(s any) error {
	return validate.Struct(s)
}

func Var(v any, validators string) error {
	return validate.Var(v, validators)
}

func JSON(j []byte, t reflect.Type) (error, any) {
	v := reflect.Zero(t).Interface()
	err := json.Unmarshal(j, &v)
	if err != nil {
		return err, nil
	}
	err = validate.Struct(v)
	return err, v
}
