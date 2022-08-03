package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the last element of slice.
//
// Complexity: O(1)
func Latest(slice interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	s := reflect.ValueOf(slice)
	if s.Len() == 0 {
		panic("slice is empty")
	}

	return s.Index(s.Len() - 1).Interface()
}

// Gets the last element of slice.
//
// Complexity: O(1)
func Last(slice interface{}) interface{} {
	return Latest(slice)
}
