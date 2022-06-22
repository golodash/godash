package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Removes all elements from slice that the passed function returns true on them
// and returns a slice of remaining elements and a slice of removed elements.
// The passed function will invoke with one argument.
//
// example for 'function':
//
//  func isOdd(n interface{}) bool {
//    return n.(int)%2 != 0
//  }
//
// Complexity: O(n)
func RemoveBy(slice interface{}, function func(interface{}) bool) (interface{}, interface{}) {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	removed := reflect.MakeSlice(reflect.TypeOf(slice), 0, sliceValue.Len())
	for i := sliceValue.Len() - 1; i >= 0; i-- {
		item := sliceValue.Index(i)
		if res := function(item.Interface()); res {
			removed = reflect.Append(removed, item)
			sliceValue = reflect.AppendSlice(sliceValue.Slice(0, i), sliceValue.Slice(i+1, sliceValue.Len()))
		}
	}

	return sliceValue.Interface(), removed.Interface()
}
