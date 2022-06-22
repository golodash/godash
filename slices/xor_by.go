package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like Xor except that it accepts a function which is
// invoked for each element of each slices for comparing them.
//
// example for 'function':
//
//  func makeInt(input interface{}) interface{} {
//    return int(input.(float64))
//  }
//
// Complexity: O(n)
//
// n = number of all elements in both 'slice1' and 'slice2'
func XorBy(slice1, slice2 interface{}, function func(interface{}) interface{}) interface{} {
	if ok := internal.SliceCheck(slice1); !ok {
		panic("passed 'slice1' variable is not slice type")
	}
	if ok := internal.SliceCheck(slice2); !ok {
		panic("passed 'slice2' variable is not slice type")
	}

	slice1Value := reflect.ValueOf(slice1)
	slice2Value := reflect.ValueOf(slice2)
	length := slice1Value.Len() + slice2Value.Len()
	var oneTimeSeenItems = reflect.Value{}
	if slice1Value.Type().String() == slice2Value.Type().String() {
		oneTimeSeenItems = reflect.MakeSlice(slice1Value.Type(), 0, length)
	} else {
		oneTimeSeenItems = reflect.MakeSlice(reflect.TypeOf([]interface{}{}), 0, length)
	}

	// find which items repeated more than once
	seenMap := map[interface{}]bool{}
	RepeatMap := map[interface{}]bool{}
	for i := 0; i < slice1Value.Len(); i++ {
		item := slice1Value.Index(i).Interface()
		compare := function(item)
		if _, ok := seenMap[compare]; !ok {
			seenMap[compare] = true
		} else {
			RepeatMap[compare] = true
		}
	}
	for i := 0; i < slice2Value.Len(); i++ {
		item := slice2Value.Index(i).Interface()
		compare := function(item)
		if _, ok := seenMap[compare]; !ok {
			seenMap[compare] = true
		} else {
			RepeatMap[compare] = true
		}
	}

	// add all items except items which appeared more than once
	for i := 0; i < slice1Value.Len(); i++ {
		item := slice1Value.Index(i).Interface()
		compare := function(item)
		if _, ok := RepeatMap[compare]; !ok {
			oneTimeSeenItems = reflect.Append(oneTimeSeenItems, reflect.ValueOf(item))
		}
	}
	for i := 0; i < slice2Value.Len(); i++ {
		item := slice2Value.Index(i).Interface()
		compare := function(item)
		if _, ok := RepeatMap[compare]; !ok {
			oneTimeSeenItems = reflect.Append(oneTimeSeenItems, reflect.ValueOf(item))
		}
	}

	return oneTimeSeenItems.Interface()
}
