package math

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Computes number rounded to precision.
//
// Complexity: O(1)
func Round(number interface{}, precision int) interface{} {
	if !internal.IsNumber(number) {
		panic("'number' is not a number")
	}

	numberValue := reflect.ValueOf(number)
	floatNumber := numberValue.Convert(reflect.TypeOf(1.0)).Float()
	tenPowered := reflect.ValueOf(Power(10, precision+1)).Convert(reflect.TypeOf(1.0)).Float()
	comparator := int(floatNumber*tenPowered) % 10
	if comparator < 0 {
		comparator = -comparator
	}
	if comparator >= 5 {
		return Ceil(number, precision)
	} else {
		return Floor(number, precision)
	}
}
