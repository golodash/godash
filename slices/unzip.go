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
func Unzip(slice interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	var output = reflect.Value{}
	sliceItemType := reflect.TypeOf(slice)
	sliceValue := reflect.ValueOf(slice)
	itItInterface := false
	for sliceItemType.Kind() == reflect.Slice {
		sliceItemType = sliceItemType.Elem()
	}
	if sliceItemType.Kind() != reflect.Interface {
		output = reflect.MakeSlice(reflect.TypeOf(slice), 0, sliceValue.Len()*2)
	} else {
		output = reflect.MakeSlice(reflect.TypeOf([][]interface{}{}), 0, sliceValue.Len()*2)
		itItInterface = true
	}

	// Check length of all slices
	length := -1
	if sliceValue.Len() != 0 {
		for i := 0; i < sliceValue.Len(); i++ {
			itemValue := reflect.ValueOf(sliceValue.Index(i).Interface())
			if err := internal.SliceCheck(sliceValue.Index(i).Interface()); err != nil {
				return nil, err
			}
			if length == -1 {
				length = itemValue.Len()
			}
			if itemValue.Len() != length {
				return nil, fmt.Errorf("item in %d index is not the same length with its previous item", i)
			}
		}
	} else {
		return output.Interface(), nil
	}

	// Actual unzip
	tempMap := map[string]interface{}{}
	for j := 0; j < length; j++ {
		indexMap := fmt.Sprint(j)
		for i := 0; i < sliceValue.Len(); i++ {
			innerSliceValue := reflect.ValueOf(sliceValue.Index(i).Interface())
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

	return output.Interface(), nil
}
