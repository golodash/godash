package numbers

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Clamps number within the inclusive lower and upper bounds.
//
// Complexity: O(1)
func Clamp(number, lower, upper interface{}) interface{} {
	if !internal.IsNumber(number) {
		panic("'number' is not a number")
	}
	if !internal.IsNumber(lower) {
		panic("'lower' is not a number")
	}
	if !internal.IsNumber(upper) {
		panic("'upper' is not a number")
	}

	numberFloat := reflect.ValueOf(number).Convert(reflect.TypeOf(1.0)).Float()
	lowerFloat := reflect.ValueOf(lower).Convert(reflect.TypeOf(1.0)).Float()
	upperFloat := reflect.ValueOf(upper).Convert(reflect.TypeOf(1.0)).Float()

	if lowerFloat > upperFloat {
		panic("'upper' has to be higher or equal to 'lower'")
	}

	outputType := internal.GetOutputNumberType(reflect.Zero(internal.GetOutputNumberType(number, lower)).Interface(), upper)

	if numberFloat <= lowerFloat {
		return reflect.ValueOf(lowerFloat).Convert(outputType).Interface()
	} else if numberFloat >= upperFloat {
		return reflect.ValueOf(upperFloat).Convert(outputType).Interface()
	} else {
		return reflect.ValueOf(numberFloat).Convert(outputType).Interface()
	}
}
