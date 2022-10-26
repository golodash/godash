package slices

import (
	"reflect"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

// Gets the index at which the first occurrence of value is found in slice
// with equality comparisons. If 'from' is negative, it's used as the offset
// from the end of slice.
//
// Complexity: O(n)
func IndexOf(slice, value interface{}, from int) int {
	return indexOf(slice, value, from, true)
}

func indexOf(slice, value interface{}, from int, ltr bool) int {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		return -1
	}

	if from < 0 {
		from = sliceValue.Len() + from
	} else if from >= sliceValue.Len() {
		panic("'from' index is out of range")
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
		if !generals.Same(sliceValue.Index(i).Interface(), value) {
			continue
		} else {
			return i
		}
	}

	return -1
}
