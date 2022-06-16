package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a slice from a slice from 'start' up to 'to', but not including, end.
//
// Complexity: O(1)
func Slice(slice interface{}, from, to int) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	if from < 0 {
		from = sliceValue.Len() + from
	}
	if to < 0 {
		to = sliceValue.Len() + to
	}

	if from > to {
		return nil, errors.New("'from' is bigger than 'to'")
	}

	if from >= sliceValue.Len() && from != 0 {
		return nil, errors.New("'from' should be in range of 'slice'")
	}
	if to > sliceValue.Len() {
		return nil, errors.New("'to' should be in range of 'slice'")
	}

	if sliceValue.Len() == 0 {
		return slice, nil
	}

	return sliceValue.Slice(from, to).Interface(), nil
}
