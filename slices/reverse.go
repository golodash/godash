package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

func Reverse(slice interface{}) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}
	length := reflect.ValueOf(slice).Len()
	swap := reflect.Swapper(slice)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return slice, nil
}
