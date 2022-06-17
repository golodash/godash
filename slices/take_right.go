package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns a sub slice of a given slice with (n) elements taken from the end.
//
// Complexity: O(1)
func TakeRight(slice interface{}, number int) (interface{}, error) {
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

	return values.Slice(values.Len()-number, values.Len()).Interface(), nil
}
