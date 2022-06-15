package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method finds first index of the given slice which the given
// function on that element, returns true.
func FindIndex(slice interface{}, function func(interface{}) bool) (int, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return -1, err
	}

	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		if function(reflect.ValueOf(sliceValue.Index(i).Interface()).Interface()) {
			return i, nil
		}
	}

	return -1, nil
}
