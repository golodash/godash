package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Fills a slice with a value from 'start' up to but not including 'end'.
//
// Complexity: O(n)
//
// n = end - start
func Fill(slice, value interface{}, start, end int) interface{} {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	outputValue := reflect.MakeSlice(reflect.TypeOf(slice), length, length)
	reflect.Copy(outputValue, sliceValue)

	if end < 0 {
		panic("negative values for 'end' variable is not accepted")
	} else if end > length {
		panic("'end' variable is bigger that slice length")
	}

	if start < 0 {
		panic("negative values for 'start' variable is not accepted")
	} else if start > end {
		panic("'start' variable is bigger than 'end' variable")
	}

	valueValue := reflect.ValueOf(value)
	for i := start; i < end; i++ {
		outputValue.Index(i).Set(valueValue)
	}

	return outputValue.Interface()
}
