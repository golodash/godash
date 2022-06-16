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
func Pull(slice, values interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}
	if err := internal.SliceCheck(values); err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	outputValue := reflect.MakeSlice(sliceValue.Type(), 0, sliceValue.Len())
	valuesValue := reflect.ValueOf(values)
	for i := 0; i < sliceValue.Len(); i++ {
		add := true
		for j := 0; j < valuesValue.Len(); j++ {
			if ok, _ := internal.Same(sliceValue.Index(i).Interface(), valuesValue.Index(j).Interface()); ok {
				add = false
				break
			}
		}

		if add {
			outputValue = reflect.Append(outputValue, sliceValue.Index(i))
		}
	}

	return outputValue.Interface(), nil
}
