package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the index at which the first occurrence of value is found in slice
// with equality comparisons. If fromIndex is negative, it's used as the offset
// from the end of slice.
//
// Note: In comparing fields of a struct, unexported fields
// are ignored.
func IndexOf(slice interface{}, value interface{}, from ...int) (int, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return -1, err
	}

	fr := 0
	sliceValue := reflect.ValueOf(slice)
	if len(from) != 0 {
		if from[0] >= 0 {
			fr = from[0]
		} else {
			fr = (sliceValue.Len() - 1) + from[0]
		}
	}

	for i := fr; i < sliceValue.Len(); i++ {
		res, err := same(sliceValue.Index(i).Interface(), value)
		if err != nil {
			return -1, err
		}

		if res {
			return i, nil
		}
	}

	return -1, nil
}
