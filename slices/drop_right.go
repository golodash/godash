package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a sub slice from a slice with n elements
// dropped from the end.
//
// By default n is equal to 0
func DropRight(slice interface{}, n ...int) ([]interface{}, error) {
	var err error = internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	s := reflect.ValueOf(slice)
	var to int = 0
	if len(n) > 0 {
		to = n[0]
	}

	if s.Len() < to {
		return nil, errors.New("num is bigger than slice length")
	}

	output, err := internal.InterfaceToSlice(s.Slice(0, s.Len()-to).Interface())
	if err != nil {
		return nil, err
	}

	return output, nil
}
