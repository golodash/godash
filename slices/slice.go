package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a slice from a slice from 'start' up to 'to', but not including, end.
//
// Complexity: O(1)
func Slice(slice interface{}, from, to int) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	if from < 0 {
		from = sliceValue.Len() + from
	}
	if to < 0 {
		to = sliceValue.Len() + to
	}

	if from > to {
		panic("'from' is bigger than 'to'")
	}

	if from >= sliceValue.Len() && from != 0 {
		panic("'from' should be in range of 'slice'")
	}
	if to > sliceValue.Len() {
		panic("'to' should be in range of 'slice'")
	}

	if sliceValue.Len() == 0 {
		return slice
	}

	return sliceValue.Slice(from, to).Interface()
}
