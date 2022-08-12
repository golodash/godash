package maths

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Check number lower than end and greater then start
//
// Complexity: O(1)
func InRange(number, start, end interface{}) interface{} {
	if !internal.IsNumber(number) {
		panic("'number' is not a number")
	}
	if !internal.IsNumber(start) {
		panic("'start' is not a number")
	}
	if !internal.IsNumber(end) {
		panic("'end' is not a number")
	}

	numberValue := reflect.ValueOf(number).Convert(reflect.TypeOf(1.0)).Float()
	startValue := reflect.ValueOf(start).Convert(reflect.TypeOf(1.0)).Float()
	endValue := reflect.ValueOf(end).Convert(reflect.TypeOf(1.0)).Float()

	if startValue <= endValue {
		if numberValue >= startValue && numberValue < endValue {
			return true
		} else {
			return false
		}
	} else {
		if numberValue >= endValue && numberValue < startValue {
			return true
		} else {
			return false
		}
	}
}
