package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the first element of slice.
func Head(slice interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
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
