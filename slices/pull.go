package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Removes all given values from slice.
func Pull(slice interface{}, values interface{}) ([]interface{}, error) {
	s, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}
	if err := internal.SliceCheck(values); err != nil {
		return nil, err
	}

	valuesValue := reflect.ValueOf(values)
	for i := 0; i < len(s); i++ {
		for j := 0; j < valuesValue.Len(); j++ {
			if ok, _ := internal.Same(s[i], valuesValue.Index(j).Interface()); ok {
				s = append(s[:i], s[i+1:]...)
				i = i - 1
				break
			}
		}
	}

	return s, nil
}
