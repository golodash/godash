package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Flattens slice 'depth' level deep.
//
// Complexity: O(n)
//
// n = count of all elements at 'depth' level of 'slice'
func FlattenDepth(slice interface{}, depth int) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	if depth == 0 {
		return slice
	} else if depth < 0 {
		return negativeFlatten(slice, -depth)
	} else {
		return recursiveFlattenDepth(slice, depth, getTypeInGivenDepth(slice, depth)).Interface()
	}
}

func getTypeInGivenDepth(slice interface{}, depth int) reflect.Type {
	sliceItemType := reflect.TypeOf(slice)
	for sliceItemType.Kind() == reflect.Slice {
		if depth == 0 {
			break
		}
		sliceItemType = sliceItemType.Elem()
		depth--
	}

	return sliceItemType
}

func negativeFlatten(slice interface{}, depth int) interface{} {
	sliceValue := reflect.ValueOf(slice)
	if depth > 0 {
		returnedSlice := negativeFlatten(slice, depth-1)
		sliceValue = reflect.Append(reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(returnedSlice)), 0, 1), reflect.ValueOf(returnedSlice))
	}

	return sliceValue.Interface()
}

func recursiveFlattenDepth(slice interface{}, depth int, itemType reflect.Type) reflect.Value {
	s := reflect.MakeSlice(reflect.SliceOf(itemType), 0, 0)
	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		item := reflect.ValueOf(sliceValue.Index(i).Interface())
		if item.Kind() == reflect.Slice && depth != 0 {
			s = reflect.AppendSlice(s, recursiveFlattenDepth(item.Interface(), depth-1, itemType))
		} else {
			s = reflect.Append(s, item)
		}
	}

	return s
}
