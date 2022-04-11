package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a sub slice from a slice with n elements
// dropped from the beginning.
//
// By default n is equal to 0
func Drop(slice interface{}, n ...int) ([]interface{}, error) {
	var err error = internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	s := reflect.ValueOf(slice)
	var from int = 0
	if len(n) > 0 {
		from = n[0]
	}

	if s.Len() < from {
		return nil, errors.New("num is bigger than slice length")
	}

	output, err := internal.InterfaceToSlice(s.Slice(from, s.Len()).Interface())
	if err != nil {
		return nil, err
	}

	return output, nil
}
