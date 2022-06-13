package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

func TakeRightWhile(slice interface{}, function func(interface{}) bool) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	i := 0
	for i = sliceValue.Len() - 1; i > -1; i-- {
		item := sliceValue.Index(i).Interface()
		if !function(item) {
			break
		}
	}

	return sliceValue.Slice(0, i+1).Interface(), nil
}
