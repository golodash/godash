package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the (n)th element of a slice. If n is negative, the (nth)
// element from the end is returned.
//
// Complexity: O(1)
func Nth(slice interface{}, index int) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	values := reflect.ValueOf(slice)
	if index < 0 {
		index = values.Len() + index
	}

	if index >= values.Len() {
		return nil, errors.New("index out of range")
	}

	if values.Len() == 0 {
		return nil, nil
	}

	return values.Index(index).Interface(), nil
}
