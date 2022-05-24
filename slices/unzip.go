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
		return nil, nil
	}

	// Actual unzip
	tempMap := map[string]interface{}{}
	for j := 0; j < length; j++ {
		tempTypeString := reflect.TypeOf(reflect.ValueOf(sliceValue.Index(0).Interface()).Index(j).Interface()).String()
		indexMap := fmt.Sprint(j)
		for i := 0; i < sliceValue.Len(); i++ {
			innerSliceValue := reflect.ValueOf(sliceValue.Index(i).Interface())
			itemValue := reflect.ValueOf(innerSliceValue.Index(j).Interface())
			if tempTypeString != itemValue.Type().String() {
				return nil, errors.New("these values types are edited and unzip can't happen")
			}

			if _, ok := tempMap[indexMap]; !ok {
				tempMap[indexMap] = reflect.MakeSlice(reflect.SliceOf(itemValue.Type()), 0, length).Interface()
			}

			tempMap[indexMap] = reflect.Append(reflect.ValueOf(tempMap[indexMap]), itemValue).Interface()
		}
	}

	// Put outputs into slice
	output := []interface{}{}
	for i := 0; i < len(tempMap); i++ {
		output = append(output, tempMap[fmt.Sprint(i)])
	}

	return output, nil
}
