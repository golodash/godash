package slices

import (
	"fmt"
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like Zip except that it accepts a slice of grouped elements
// and creates a slice regrouping the elements to their pre-zip configuration.
//
// Complexity: O(n)
func Unzip(slices interface{}) interface{} {
	if !internal.SliceCheck(slices) {
		panic("passed 'slices' variable is not slice type")
	}

	var output = reflect.Value{}
	sliceItemType := reflect.TypeOf(slices)
	slicesValue := reflect.ValueOf(slices)
	itItInterface := false
	for sliceItemType.Kind() == reflect.Slice {
		sliceItemType = sliceItemType.Elem()
	}
	if sliceItemType.Kind() != reflect.Interface {
		output = reflect.MakeSlice(reflect.TypeOf(slices), 0, slicesValue.Len()*2)
	} else {
		output = reflect.MakeSlice(reflect.TypeOf([][]interface{}{}), 0, slicesValue.Len()*2)
		itItInterface = true
	}

	// Check length of all slices
	length := -1
	if slicesValue.Len() != 0 {
		for i := 0; i < slicesValue.Len(); i++ {
			itemValue := reflect.ValueOf(slicesValue.Index(i).Interface())
			if !internal.SliceCheck(slicesValue.Index(i).Interface()) {
				panic(fmt.Sprintf("item in %d index is not a slice", i))
			}
			if length == -1 {
				length = itemValue.Len()
			}
			if itemValue.Len() != length {
				panic(fmt.Sprintf("item in %d index is not the same length with its previous item", i))
			}
		}
	} else {
		return output.Interface()
	}

	// Actual unzip
	tempMap := map[string]interface{}{}
	for j := 0; j < length; j++ {
		indexMap := fmt.Sprint(j)
		for i := 0; i < slicesValue.Len(); i++ {
			innerSliceValue := reflect.ValueOf(slicesValue.Index(i).Interface())
			itemValue := reflect.ValueOf(innerSliceValue.Index(j).Interface())

			if _, ok := tempMap[indexMap]; !ok {
				if itItInterface {
					tempMap[indexMap] = reflect.MakeSlice(reflect.TypeOf([]interface{}{}), 0, length).Interface()
				} else {
					tempMap[indexMap] = reflect.MakeSlice(reflect.SliceOf(itemValue.Type()), 0, length).Interface()
				}
			}

			tempMap[indexMap] = reflect.Append(reflect.ValueOf(tempMap[indexMap]), itemValue).Interface()
		}
	}

	// Put outputs into slice
	for i := 0; i < len(tempMap); i++ {
		output = reflect.Append(output, reflect.ValueOf(tempMap[fmt.Sprint(i)]))
	}

	return output.Interface()
}
