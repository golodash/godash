package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns a slice of unique values that is the symmetric difference of the given slices.
//
// Complexity: O(n)
//
// n = number of all elements in both 'slice1' and 'slice2'
func Xor(slice1, slice2 interface{}) (interface{}, error) {
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
		if _, ok := seenMap[item]; !ok {
			seenMap[item] = true
		} else {
			RepeatMap[item] = true
		}
	}
	for i := 0; i < slice2Value.Len(); i++ {
		item := slice2Value.Index(i).Interface()
		if _, ok := seenMap[item]; !ok {
			seenMap[item] = true
		} else {
			RepeatMap[item] = true
		}
	}

	// add all items except items which appeared more than once
	for i := 0; i < slice1Value.Len(); i++ {
		item := slice1Value.Index(i).Interface()
		if _, ok := RepeatMap[item]; !ok {
			oneTimeSeenItems = reflect.Append(oneTimeSeenItems, reflect.ValueOf(item))
		}
	}
	for i := 0; i < slice2Value.Len(); i++ {
		item := slice2Value.Index(i).Interface()
		if _, ok := RepeatMap[item]; !ok {
			oneTimeSeenItems = reflect.Append(oneTimeSeenItems, reflect.ValueOf(item))
		}
	}

	return oneTimeSeenItems.Interface(), nil
}
