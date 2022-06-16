package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like Intersection except that it accepts comparator which is
// invoked to compare elements of slices. The order and references of result
// values are determined by the first slice. The comparator is invoked with
// two arguments: (slice, otherSlice).
func IntersectionBy(slices interface{}, function func(interface{}, interface{}) bool) (interface{}, error) {
	if err := internal.SliceCheck(slices); err != nil {
		return nil, err
	}

	sliceItemType := reflect.TypeOf(slices)
	if sliceItemType = sliceItemType.Elem(); sliceItemType.Kind() == reflect.Slice {
		sliceItemType = sliceItemType.Elem()
	}

	sliceValue := reflect.ValueOf(slices)
	length := 0
	for i := 0; i < sliceValue.Len(); i++ {
		subSlice := reflect.ValueOf(sliceValue.Index(i).Interface())
		if err := internal.SliceCheck(subSlice); err != nil {
			continue
		}

		length += subSlice.Len()
	}

	outputSlice := reflect.MakeSlice(reflect.SliceOf(sliceItemType), 0, length)
	for i := 0; i < sliceValue.Len(); i++ {
		subSlice := reflect.ValueOf(sliceValue.Index(i).Interface())
		if err := internal.SliceCheck(subSlice.Interface()); err != nil {
			continue
		}

		for j := i + 1; j < sliceValue.Len(); j++ {
			secondSubSlice := reflect.ValueOf(sliceValue.Index(i + 1).Interface())
			if function(subSlice.Interface(), secondSubSlice.Interface()) {
				outputSlice = reflect.AppendSlice(outputSlice, subSlice)
				break
			}
		}
	}

	return outputSlice.Interface(), nil
}
