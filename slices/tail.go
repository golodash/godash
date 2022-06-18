package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns all but the first element of slice.
//
// Complexity: O(1)
func Tail(slice interface{}) (interface{}, error) {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		return slice, nil
	}

	return sliceValue.Slice(1, sliceValue.Len()).Interface(), nil
}
