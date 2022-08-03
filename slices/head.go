package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the first element of slice.
//
// Complexity: O(1)
func Head(slice interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		panic("slice is empty")
	}

	return sliceValue.Index(0).Interface()
}
