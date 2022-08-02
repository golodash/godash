package math

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns input to the power of number.
//
// O(number)
func Power(input interface{}, number int) interface{} {
	if !internal.IsNumber(input) {
		panic("'input' is not a number")
	}

	if number == 0 {
		return 1
	}

	inputFloat := reflect.ValueOf(input).Convert(reflect.TypeOf(0.1)).Float()
	inputType := reflect.TypeOf(input)

	isNegative := false
	if inputFloat < 0 {
		isNegative = true
		inputFloat = -inputFloat
	}
	isNegativeNumber := false
	if number < 0 {
		isNegativeNumber = true
		number = -number
	}

	output := inputFloat
	for i := 1; i < number; i++ {
		output *= inputFloat
	}

	if isNegative && number%2 != 0 {
		output = -output
	}
	if isNegativeNumber {
		output = 1.0 / output
		if isFloat(output) {
			return output
		}
	}

	return reflect.ValueOf(output).Convert(inputType).Interface()
}

func isFloat(number interface{}) bool {
	switch reflect.TypeOf(number).Kind() {
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	}
	return false
}
