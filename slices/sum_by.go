package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like Sum except that it accepts a function
// which is invoked for each element in slice to generate the
// value to be summed.
//
// example for 'function':
//
//	type myObject struct {
//		rank int
//	}
//
//	func returnRank(value1 interface{}) interface{} {
//		return value1.(myObject).rank
//	}
//
// Complexity: O(n)
func SumBy(slice interface{}, function func(interface{}) interface{}) interface{} {
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
		element := reflect.ValueOf(function(sliceValue.Index(i).Interface()))

		if internal.CanFloat(element.Interface()) {
			sum = reflect.ValueOf(sum.Float() + element.Float())
		} else {
			sum = reflect.ValueOf(sum.Float() + element.Convert(floatType).Float())
		}
	}

	return sum.Convert(sliceElementType).Interface()
}
