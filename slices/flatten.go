package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Flattens slice a single level deep.
//
// Complexity: O(n)
func Flatten(slice interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	var output reflect.Value
	if sliceValue.Type().Elem().Kind() == reflect.Interface {
		output = reflect.MakeSlice(reflect.TypeOf([]interface{}{}), 0, sliceValue.Len())
	} else if sliceValue.Type().Elem().Kind() == reflect.Slice {
		output = reflect.MakeSlice(sliceValue.Type().Elem(), 0, sliceValue.Len())
	} else {
		return slice
	}

	for i := 0; i < sliceValue.Len(); i++ {
		item := reflect.ValueOf(sliceValue.Index(i).Interface())
		if item.Kind() == reflect.Slice {
			for j := 0; j < item.Len(); j++ {
				output = reflect.Append(output, item.Index(j))
			}
		} else {
			output = reflect.Append(output, item)
		}
	}

	return output.Interface()
}
