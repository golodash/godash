package maths

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Check if number is lower than 'upper' and greater than 'lower'.
//
// Complexity: O(1)
func InRange(number, lower, upper interface{}) interface{} {
	if !internal.IsNumber(number) {
		panic("'number' is not a number")
	}
	if !internal.IsNumber(lower) {
		panic("'start' is not a number")
	}
	if !internal.IsNumber(upper) {
		panic("'end' is not a number")
	}

	numberFloat := reflect.ValueOf(number).Convert(reflect.TypeOf(1.0)).Float()
	lowerFloat := reflect.ValueOf(lower).Convert(reflect.TypeOf(1.0)).Float()
	upperFloat := reflect.ValueOf(upper).Convert(reflect.TypeOf(1.0)).Float()

	if lowerFloat > upperFloat {
		panic("'upper' has to be higher or equal to 'lower'")
	}

	if numberFloat >= lowerFloat && numberFloat <= upperFloat {
		return true
	} else {
		return false
	}
}
