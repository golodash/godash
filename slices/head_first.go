package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the first element of slice.
func Head(slice interface{}) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	s := reflect.ValueOf(slice)

	if s.Len() == 0 {
		return nil, errors.New("slice is empty")
	}

	return s.Index(0).Interface(), nil
}

// Gets the first element of slice.
func First(slice interface{}) (interface{}, error) {
	return Head(slice)
}
