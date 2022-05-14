package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

//Creates a sub slice of a given slice with (n) elements taken from the beginning.
func Take(slice interface{}, number int) (interface{}, error) {
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

	newSlice := reflect.MakeSlice(reflect.TypeOf(slice), 0, (values.Len() - number))
	for i := 0; i < number; i++ {
		newSlice = reflect.Append(newSlice, values.Index(i))
	}

	return newSlice.Interface(), nil
}
