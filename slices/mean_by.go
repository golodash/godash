package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like Mean except that it accepts a function
// which is invoked for each element in slice to generate the
// value to be averaged.
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
func MeanBy(slice interface{}, function func(interface{}) interface{}) float64 {
	if !internal.SliceCheck(slice) {
		panic("'slice' is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Len() == 0 {
		return 0
	}

	floatType := reflect.TypeOf(1.0)
	sum := 0.0
	for i := 0; i < sliceValue.Len(); i++ {
		sum += reflect.ValueOf(function(sliceValue.Index(i).Interface())).Convert(floatType).Float()
	}

	return sum / float64(sliceValue.Len())
}
