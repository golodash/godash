package slices

import (
	"errors"
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
	sliceValue := reflect.ValueOf(slice)
	fr := 0

	if len(from) != 0 {
		if from[0] >= 0 {
			fr = from[0]
		} else {
			fr = (sliceValue.Len() - 1) + from[0]
		}
	}
	if fr >= sliceValue.Len() {
		return -1, errors.New("'from' index is out of range")
	} else if fr <= -1 {
		return -1, nil
	}
	return indexOf(slice, value, fr, true)
}

func indexOf(slice interface{}, value interface{}, from int, ltr bool) (int, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return -1, err
	}

	sliceValue := reflect.ValueOf(slice)

	var until int
	var count int
	if ltr {
		until = sliceValue.Len()
		count = +1
	} else {
		until = -1
		if sliceValue.Len() == 0 {
			until = 0
		}
		count = -1
	}

	compare := func(i, until int) bool {
		if until == -1 {
			return i > until
		} else {
			return i < until
		}
	}

	for i := from; compare(i, until); i += count {
		res, err := internal.Same(sliceValue.Index(i).Interface(), value)
		if err != nil {
			return -1, err
		}

		if res {
			return i, nil
		}
	}

	return -1, nil
}
