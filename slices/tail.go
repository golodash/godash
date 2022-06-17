package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns all but the first element of slice.
//
// Complexity: O(1)
func Tail(slice interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		return slice, nil
	}

	return sliceValue.Slice(1, sliceValue.Len()).Interface(), nil
}
