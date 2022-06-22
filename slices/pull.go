package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Removes all given values from slice.
//
// Complexity: O(n*m)
//
// n = length of 'slice'
//
// m = length of 'values'
func Pull(slice, values interface{}) interface{} {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}
	if ok := internal.SliceCheck(values); !ok {
		panic("passed 'values' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	outputValue := reflect.MakeSlice(sliceValue.Type(), 0, sliceValue.Len())
	valuesValue := reflect.ValueOf(values)
	for i := 0; i < sliceValue.Len(); i++ {
		add := true
		for j := 0; j < valuesValue.Len(); j++ {
			if ok := internal.Same(sliceValue.Index(i).Interface(), valuesValue.Index(j).Interface()); ok {
				add = false
				break
			}
		}

		if add {
			outputValue = reflect.Append(outputValue, sliceValue.Index(i))
		}
	}

	return outputValue.Interface()
}
