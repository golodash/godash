package slices

import (
	"errors"
	"reflect"
)

// This method is like IndexOf except that it iterates over elements of
// slice from right to left.
func LastIndexOf(slice interface{}, value interface{}, from ...int) (int, error) {
	sliceValue := reflect.ValueOf(slice)
	fr := sliceValue.Len() - 1

	if len(from) != 0 {
		if from[0] >= 0 {
			fr = from[0]
		} else {
			fr = sliceValue.Len() + from[0]
		}
	}
	if fr >= sliceValue.Len() {
		return -1, errors.New("`from` index is out of range")
	} else if fr <= -1 {
		return -1, nil
	}
	return indexOf(slice, value, fr, false)
}
