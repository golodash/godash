package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets all but the last element of slice.
//
// Complexity: O(1)
func Initial(slice interface{}) (interface{}, error) {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		return slice, nil
	}

	return sliceValue.Slice(0, sliceValue.Len()-1).Interface(), nil
}
