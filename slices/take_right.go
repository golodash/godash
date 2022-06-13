package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

//Returns a sub slice of a given slice with (n) elements taken from the end.
func TakeRight(slice interface{}, number int) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}
	values := reflect.ValueOf(slice)
	if (number > values.Len() && values.Len() != 0) || (number < 0 && values.Len() != 0) {
		return nil, errors.New("'number' should be in range of slice's length")
	}
	if values.Len() == 0 {
		return slice, nil
	}
	j := values.Len() - number
	return values.Slice(j, values.Len()).Interface(), nil
}
