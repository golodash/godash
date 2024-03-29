package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets all but the last element of slice.
//
// Complexity: O(1)
func Initial(slice interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		return slice
	}

	return sliceValue.Slice(0, sliceValue.Len()-1).Interface()
}
