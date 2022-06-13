package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

func TakeWhile(slice interface{}, function func(interface{}) bool) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	i := 0
	for i = 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		if !function(item) {
			break
		}
	}

	return sliceValue.Slice(i, sliceValue.Len()).Interface(), nil
}
