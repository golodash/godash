package math

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Computes number rounded up to precision.
//
// Complexity: O(1)
func Ceil(input interface{}, precision int) interface{} {
	if !internal.IsNumber(input) {
		panic("'input' is not a number")
	}

	inputValue := reflect.ValueOf(input)
	inputType := inputValue.Type()

	isNegative := false
	if internal.CanInt(input) && inputValue.Int() < 0 {
		isNegative = true
		inputValue = reflect.ValueOf(-inputValue.Int())
	} else if internal.CanFloat(input) && inputValue.Float() < 0 {
		isNegative = true
		inputValue = reflect.ValueOf(-inputValue.Float())
	}

	if precision > 0 {
		floatValue := inputValue.Convert(reflect.TypeOf(0.1)).Float()
		tenPowered := float64(Power(10, precision).(int))
		floatValuePowered := floatValue * tenPowered
		output := float64(int(floatValuePowered)+1) / tenPowered

		if isNegative {
			output = -output
		}
		return reflect.ValueOf(output).Convert(inputType).Interface()
	} else if precision < 0 {
		precision = -precision
		intValue := inputValue.Convert(reflect.TypeOf(0)).Int()
		output := intValue + int64(Power(10, precision).(int)) - intValue%int64((Power(10, precision).(int)))

		if isNegative {
			output = -output
		}
		return reflect.ValueOf(output).Convert(inputType).Interface()
	} else {
		output := inputValue.Convert(reflect.TypeOf(0)).Int() + int64(1)

		if isNegative {
			output = -output
		}
		return reflect.ValueOf(output).Convert(inputType).Interface()
	}
}
