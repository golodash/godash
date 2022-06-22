package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like Intersection except that it accepts comparator which is
// invoked to compare elements of 'slices'. The order and references of result
// values are determined by the first slice. The comparator is invoked with
// two arguments: (slice1, slice2).
//
// example for function:
//
//  func isEqual(value1, value2 interface{}) bool {
//    v1 := value1.([]int)
//    v2 := value2.([]int)
//    same := true
//    for i := 0; i < v1.len(); i ++ {
//      if v1[i] != v2[i] {
//        same = false
//        break
//      }
//    }
//    return same
//  }
//
// Complexity: O(n*log(n))
func IntersectionBy(slices interface{}, function func(interface{}, interface{}) bool) interface{} {
	if ok := internal.SliceCheck(slices); !ok {
		panic("passed 'slices' variable is not slice type")
	}

	sliceItemType := reflect.TypeOf(slices)
	if sliceItemType = sliceItemType.Elem(); sliceItemType.Kind() == reflect.Slice {
		sliceItemType = sliceItemType.Elem()
	}

	sliceValue := reflect.ValueOf(slices)
	length := 0
	for i := 0; i < sliceValue.Len(); i++ {
		subSlice := reflect.ValueOf(sliceValue.Index(i).Interface())
		if ok := internal.SliceCheck(subSlice); !ok {
			continue
		}

		length += subSlice.Len()
	}

	outputSlice := reflect.MakeSlice(reflect.SliceOf(sliceItemType), 0, length)
	for i := 0; i < sliceValue.Len(); i++ {
		subSlice := reflect.ValueOf(sliceValue.Index(i).Interface())
		if ok := internal.SliceCheck(subSlice.Interface()); !ok {
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

	return outputSlice.Interface()
}
