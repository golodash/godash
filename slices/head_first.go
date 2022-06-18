package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the first element of slice.
//
// Complexity: O(1)
func Head(slice interface{}) (interface{}, error) {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		return nil, errors.New("slice is empty")
	}

	return sliceValue.Index(0).Interface(), nil
}

// Gets the first element of slice.
func First(slice interface{}) (interface{}, error) {
	return Head(slice)
}
