package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns all but the first element of slice.
//
// Complexity: O(1)
func Tail(slice interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		return slice
	}

	return sliceValue.Slice(1, sliceValue.Len()).Interface()
}
