package maths

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns input to the power of number.
//
// Complexity: O(number)
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
		if output != float64(int(output)) {
			return output
		}
	}

	return reflect.ValueOf(output).Convert(inputType).Interface()
}
