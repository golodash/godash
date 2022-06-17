package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a sub slice of a given slice with (n) elements taken from the beginning.
//
// Complexity: O(1)
func Take(slice interface{}, number int) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	values := reflect.ValueOf(slice)
	if values.Len() == 0 {
		return slice, nil
	}
	if number > values.Len() || number < 0 {
		return nil, errors.New("'number' should be in range of slice's length")
	}

	return values.Slice(0, number).Interface(), nil
}
