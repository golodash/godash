package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a slice from a slice from `start` up `to`, but not including, end.
func Slice(slice interface{}, from, to int) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	if from > to {
		return nil, errors.New("invalid index numbers")
	}

	sliceValue := reflect.ValueOf(slice)
	if from < 0 || from >= sliceValue.Len() {
		return nil, errors.New("`from` should be in range of `slice`")
	}
	if to < 1 || to > sliceValue.Len() {
		return nil, errors.New("`to` should be in range of `slice`")
	}

	if sliceValue.Len() == 0 {
		return []interface{}{}, nil
	}

	var newSlice []interface{}
	for i := from; i < to; i++ {
		newSlice = append(newSlice, sliceValue.Index(i).Interface())
	}

	return newSlice, nil
}
