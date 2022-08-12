package maths

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

	numberValue := reflect.ValueOf(number).Convert(reflect.TypeOf(1.0)).Float()
	lowerValue := reflect.ValueOf(lower).Convert(reflect.TypeOf(1.0)).Float()
	upperValue := reflect.ValueOf(upper).Convert(reflect.TypeOf(1.0)).Float()

	if lowerValue <= upperValue {
		if numberValue <= lowerValue {
			return lowerValue
		} else if numberValue >= upperValue {
			return upperValue
		} else {
			return numberValue
		}
	} else {
		if numberValue <= upperValue {
			return upperValue
		} else if numberValue >= lowerValue {
			return lowerValue
		} else {
			return numberValue
		}
	}
}
