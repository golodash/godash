package slices

import (
	"fmt"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns a slice of grouped elements, the first of
// which contains the first elements of the given slices,
// the second of which contains the second elements of
// the given slices, and so on.
func Zip(slices interface{}) ([]interface{}, error) {
	if err := internal.SliceCheck(slices); err != nil {
		return nil, err
	}

	// Check type and length of all elements
	slicesValue := reflect.ValueOf(slices)
	length := -1
	if slicesValue.Len() != 0 {
		for i := 0; i < slicesValue.Len(); i++ {
			itemValue := reflect.ValueOf(slicesValue.Index(i).Interface())
			if err := internal.SliceCheck(slicesValue.Index(i).Interface()); err != nil {
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

	// Actual zip
	tempMap := map[string][]interface{}{}
	for j := 0; j < length; j++ {
		indexMap := fmt.Sprint(j)
		for i := 0; i < slicesValue.Len(); i++ {
			sliceValue := reflect.ValueOf(slicesValue.Index(i).Interface())
			itemValue := reflect.ValueOf(sliceValue.Index(j).Interface())
			if _, ok := tempMap[indexMap]; !ok {
				tempMap[indexMap] = []interface{}{}
			}

			tempMap[indexMap] = reflect.Append(reflect.ValueOf(tempMap[indexMap]), itemValue).Interface().([]interface{})
		}
	}

	// Put outputs into slice
	output := []interface{}{}
	for i := 0; i < len(tempMap); i++ {
		output = append(output, tempMap[fmt.Sprint(i)])
	}

	return output, nil
}
