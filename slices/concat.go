package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a new slice concatenating slice with other one.
//
// Complexity: O(n)
//
// n = number of 'values' length
func Concat(slice, values interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	valuesValue := reflect.ValueOf(values)
	sliceValue := reflect.ValueOf(slice)
	sliceType := reflect.TypeOf(slice)
	for i := 0; i < valuesValue.Len(); i++ {
		item := reflect.ValueOf(valuesValue.Index(i).Interface())
		if !item.IsValid() {
			continue
		}
		if item.Kind() == reflect.Slice {
			if sliceType.Kind() == item.Kind() || sliceType.Elem().Kind() == reflect.Interface || item.Elem().Kind() == reflect.Interface {
				for j := 0; j < item.Len(); j++ {
					innerItem := reflect.ValueOf(item.Index(j).Interface())
					if !innerItem.IsValid() {
						continue
					}
					if innerItem.Kind() == sliceType.Elem().Kind() || sliceType.Elem().Kind() == reflect.Interface {
						sliceValue = reflect.Append(sliceValue, innerItem)
					}
				}
			}
		} else {
			if item.Kind() == sliceType.Elem().Kind() || sliceType.Elem().Kind() == reflect.Interface {
				sliceValue = reflect.Append(sliceValue, item)
			}
		}
	}

	return sliceValue.Interface()
}
