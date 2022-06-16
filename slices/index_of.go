package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the index at which the first occurrence of value is found in slice
// with equality comparisons. If 'from' is negative, it's used as the offset
// from the end of slice.
func IndexOf(slice, value interface{}, from int) (int, error) {
	return indexOf(slice, value, from, true)
}

func indexOf(slice, value interface{}, from int, ltr bool) (int, error) {
	sliceValue := reflect.ValueOf(slice)
	if err := internal.SliceCheck(slice); err != nil {
		return -1, err
	}

	if from < 0 {
		from = (sliceValue.Len() - 1) + from
	} else if from >= sliceValue.Len() {
		return -1, errors.New("'from' index is out of range")
	}

	var until int
	var count int
	if ltr {
		until = sliceValue.Len()
		count = +1
	} else {
		until = -1
		count = -1
		if sliceValue.Len() == 0 {
			until = 0
		}
	}

	compare := func(i, until int) bool {
		if until == -1 {
			return i > until
		} else {
			return i < until
		}
	}

	for i := from; compare(i, until); i += count {
		if ok, err := internal.Same(sliceValue.Index(i).Interface(), value); !ok || err != nil {
			continue
		} else {
			return i, nil
		}
	}

	return -1, nil
}
