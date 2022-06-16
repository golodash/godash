package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets all but the last element of slice.
func Initial(slice interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		return slice, nil
	}

	return sliceValue.Slice(0, sliceValue.Len()-1).Interface(), nil
}
