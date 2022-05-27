// our validators are here
package validator

import (
	"errors"
	"reflect"

	"github.com/mhkarimi1383/goAPIBaseProject/types"
)

type comparableNumber interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 |
		float32 | float64 | uint | uintptr | int
}

func checkMinMax[V comparableNumber](min, max, value V) error {
	if min > max {
		return errors.New("min is bigger than max")
	}
	if min != 0 && value < min {
		return errors.New("value is less than min")
	}
	if max != 0 && value > max {
		return errors.New("value is greater than max")
	}
	return nil
}

func Validate[R types.Response, V comparableNumber](v R) error {
	typeOfValue := reflect.TypeOf(v)
	// valueOfValue := reflect.ValueOf(v).Elem()
	comparableNumberType := reflect.TypeOf((*V)(nil)).Elem()
	for i := 0; i < typeOfValue.NumField(); i++ {
		// valueField := valueOfValue.Field(i)
		if typeOfValue.Field(i).Type.Implements(comparableNumberType) {
			// CANNOT GET VALUES AS comparableNumber TYPE
			// minStr := typeOfValue.Field(i).Tag.Get("min")
			// maxStr := typeOfValue.Field(i).Tag.Get("max")
			// min, err := strconv.Atoi(minStr)
			// if err != nil {
			// 	return err
			// }
			// max, err := strconv.Atoi(maxStr)
			// if err != nil {
			// 	return err
			// }
			// err = checkMinMax(min, max, valueField.Interface().(V))
			// if err != nil {
			// 	return err
			// }
		}
	}
	return nil
}
