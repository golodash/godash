package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

//Creates a slice of array from start up to, but not including, end.
func Slice(array interface{}, from int, to int) (interface{}, error) {
	err := internal.SliceCheck(array)
	if err != nil {
		return nil, err
	}
	if from > to {
		return nil, errors.New("invalid index numbers")
	}
	values := reflect.ValueOf(array)
	if values.Len() == 0 {
		return []interface{}{}, nil
	}
	if from < 0 || from > values.Len() {
		return nil, errors.New("'from' should be in range of slice")
	}
	if to < 0 || to > values.Len() {
		return nil, errors.New("'to' should be in range of slice")
	}
	var newSlice []interface{}
	for i := from; i < to; i++ {
		newSlice = append(newSlice, values.Index(i).Interface())
	}
	return newSlice, nil
}
