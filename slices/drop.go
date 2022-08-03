package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a sub slice from passed slice with n elements dropped from the beginning.
//
// Complexity: O(1)
func Drop(slice interface{}, n int) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() < n {
		panic("'num' is bigger than slice length")
	}
	if n < 0 {
		panic("'num' is lower that zero")
	}

	return sliceValue.Slice(n, sliceValue.Len()).Interface()
}
