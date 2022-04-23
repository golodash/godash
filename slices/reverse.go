package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

func Reverse(slice interface{}) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, errors.New("the entry should be slice")
	}
	n := reflect.ValueOf(slice).Len()
	if n == 0 {
		return nil, nil
	}
	swap := reflect.Swapper(slice)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return slice, nil
}
