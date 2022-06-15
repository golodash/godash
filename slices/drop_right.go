package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a sub slice from a slice with n elements
// dropped from the end.
func DropRight(slice interface{}, n int) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() < n {
		return nil, errors.New("'num' is bigger than slice length")
	}
	if n < 0 {
		return nil, errors.New("'num' is lower that zero")
	}

	return sliceValue.Slice(0, sliceValue.Len()-n).Interface(), nil
}
