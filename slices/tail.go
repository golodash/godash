package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

func Tail(slice interface{}) ([]interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		return []interface{}{}, nil
	}

	output, err := internal.InterfaceToSlice(sliceValue.Slice(1, sliceValue.Len()).Interface())
	if err != nil {
		return nil, err
	}

	return output, nil
}
