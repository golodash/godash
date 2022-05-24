package slices

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like Zip except that it accepts
// a slice of grouped elements and creates a slice
// regrouping the elements to their pre-zip configuration.
func Unzip(slice interface{}) ([]interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	// Check length of all slices
	sliceValue := reflect.ValueOf(slice)
	length := -1
	if sliceValue.Len() != 0 {
		for i := 0; i < sliceValue.Len(); i++ {
			if sliceValue.Index(i).Kind() != reflect.Slice {
				return nil, fmt.Errorf("item in %d index is not slice type", i)
			}
			if length == -1 {
				length = sliceValue.Index(i).Len()
			}
			if sliceValue.Index(i).Len() != length {
				return nil, fmt.Errorf("item in %d index is not the same length with its previous item", i)
			}
		}
	} else {
		return nil, nil
	}

	// Actual unzip
	tempMap := map[string]interface{}{}
	for j := 0; j < length; j++ {
		tempTypeString := reflect.TypeOf(sliceValue.Index(0).Index(j).Interface()).String()
		indexMap := fmt.Sprint(j)
		for i := 0; i < sliceValue.Len(); i++ {
			if tempTypeString != reflect.TypeOf(sliceValue.Index(i).Index(j).Interface()).String() {
				return nil, errors.New("these values are edited and unzip can't happen")
			}

			if _, ok := tempMap[indexMap]; !ok {
				tempMap[indexMap] = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(sliceValue.Index(0).Index(j).Interface())), 0, length).Interface()
			}

			tempMap[indexMap] = reflect.Append(reflect.ValueOf(tempMap[indexMap]), reflect.ValueOf(sliceValue.Index(i).Index(j).Interface())).Interface()
		}
	}

	// Put outputs into slice
	output := []interface{}{}
	for i := 0; i < len(tempMap); i++ {
		output = append(output, tempMap[fmt.Sprint(i)])
	}

	return output, nil
}
