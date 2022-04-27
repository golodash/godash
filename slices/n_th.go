package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the (n)th element of a slice. If n is negative, the (nth)
// element from the end is returned.
func Nth(slice interface{}, index int) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}
	values := reflect.ValueOf(slice)
	if index >= values.Len() {
		return nil, errors.New("index out of range")
	}
	if index < 0 {
		index = values.Len() + index
	}
	if values.Len() == 0 {
		return nil, nil
	}

	return values.Index(index).Interface(), nil
}
