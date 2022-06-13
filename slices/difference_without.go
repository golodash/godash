package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns a slice of slice values not included in the
// other given slice using equality comparisons.
//
// Note: In comparing fields of a struct, unexported fields
// are ignored.
func Difference(slice interface{}, notIncluded interface{}) ([]interface{}, error) {
	if err := internal.CheckSameType(slice, notIncluded); err != nil {
		return nil, err
	}
	if err1, err2 := internal.SliceCheck(slice), internal.SliceCheck(notIncluded); err1 != nil || err2 != nil {
		if err2 != nil {
			return nil, err2
		}
		return nil, err1
	}

	s, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}

	notIn := reflect.ValueOf(notIncluded)

	for i := len(s) - 1; i > -1; i-- {
		if i >= len(s) {
			continue
		}
	firstLoop:
		for j := 0; j < notIn.Len(); j++ {
			res, err := internal.Same(s[i], notIn.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			if res {
				if i != 0 && i+1 < len(s) {
					s = append(s[0:i], s[i+1:]...)
				} else if i == 0 {
					s = s[i+1:]
				} else if i+1 >= len(s) {
					s = s[0:i]
				}
				i++
				break firstLoop
			}
		}
	}

	return s, nil
}

func Without(slice interface{}, notIncluded interface{}) ([]interface{}, error) {
	return Difference(slice, notIncluded)
}
