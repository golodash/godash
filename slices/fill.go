package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Fills a slice with a value from `start` up to but not including `end`
//
// If end == -1, iterate goes from start to end
func Fill(slice interface{}, value interface{}, start, end int) error {
	err := internal.SliceCheck(slice)
	if err != nil {
		return err
	}

	sliceValue := reflect.ValueOf(slice)
	reflectValue := reflect.ValueOf(value)

	if end == -1 {
		end = sliceValue.Len()
	} else if end < -1 {
		return errors.New("negative values for `end` variable(except -1) is not accepted")
	} else if end > sliceValue.Len() {
		end = sliceValue.Len()
	}

	if start < 0 {
		return errors.New("negative values for `start` variable is not accepted")
	} else if start > end {
		return errors.New("`start` variable is bigger than `end` variable")
	}

	for i := start; i < end; i++ {
		sliceValue.Index(i).Set(reflectValue)
	}

	return nil
}
