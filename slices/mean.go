package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Computes the mean of the values in slice.
//
// Complexity: O(n)
func Mean(slice interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("'slice' is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	sliceElementType := sliceValue.Type().Elem()

	if sliceElementType.Kind() == reflect.Interface {
		sliceElementType = reflect.TypeOf(1.0)
	}

	if sliceValue.Len() == 0 {
		return reflect.Zero(sliceElementType).Interface()
	}

	floatType := reflect.TypeOf(1.0)
	sum := reflect.Zero(floatType)
	for i := 0; i < sliceValue.Len(); i++ {
		element := reflect.ValueOf(sliceValue.Index(i).Interface())
		if internal.CanFloat(element.Interface()) {
			sum = reflect.ValueOf(sum.Float() + element.Float())
		} else {
			sum = reflect.ValueOf(sum.Float() + element.Convert(floatType).Float())
		}
	}

	return sum.Convert(sliceElementType).Interface()
}
