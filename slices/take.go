package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a sub slice of a given slice with (n) elements taken from the beginning.
//
// Complexity: O(1)
func Take(slice interface{}, number int) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	values := reflect.ValueOf(slice)
	if values.Len() == 0 {
		return slice
	}
	if number > values.Len() || number < 0 {
		panic("'number' should be in range of slice's length")
	}

	return values.Slice(0, number).Interface()
}
