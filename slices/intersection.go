package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a slice of unique values that are included in all given slices
// for equality comparisons. The order and references of result values are
// determined by the first slice.
func Intersection(slice interface{}) ([]interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	values := []interface{}{}
	doubleSeen := []interface{}{}
	for i := 0; i < sliceValue.Len(); i++ {
		item, err := internal.InterfaceToSlice((sliceValue.Index(i).Interface()))
		var ok = false
		if err != nil {
			item, ok = sliceValue.Index(i).Interface().([]interface{})
			if !ok {
				continue
			}
		}

		for j := 0; j < len(item); j++ {
			val := item[j]
			seen := false
			for k := 0; k < len(values); k++ {
				res, err := same(val, values[k])
				if err == nil && res {
					seen = true
					break
				}
			}
			if !seen {
				values = append(values, val)
			} else {
				seen = false
				for k := 0; k < len(doubleSeen); k++ {
					res, err := same(val, doubleSeen[k])
					if err == nil && res {
						seen = true
						break
					}
				}
				if !seen {
					doubleSeen = append(doubleSeen, val)
				}
			}
		}
	}

	return doubleSeen, nil
}
