package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the last element of slice.
//
// Complexity: O(1)
func Latest(slice interface{}) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	s := reflect.ValueOf(slice)
	if s.Len() == 0 {
		return nil, errors.New("slice is empty")
	}

	return s.Index(s.Len() - 1).Interface(), nil
}

// Gets the last element of slice.
//
// Complexity: O(1)
func Last(slice interface{}) (interface{}, error) {
	return Latest(slice)
}
