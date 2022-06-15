package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Fills a slice with a value from 'start' up to but not including 'end'
func Fill(slice, value interface{}, start, end int) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	outputValue := reflect.MakeSlice(reflect.TypeOf(slice), length, length)
	reflect.Copy(outputValue, sliceValue)

	if end < 0 {
		return nil, errors.New("negative values for 'end' variable is not accepted")
	} else if end > length {
		return nil, errors.New("'end' variable is bigger that slice length")
	}

	if start < 0 {
		return nil, errors.New("negative values for 'start' variable is not accepted")
	} else if start > end {
		return nil, errors.New("'start' variable is bigger than 'end' variable")
	}

	valueValue := reflect.ValueOf(value)
	for i := start; i < end; i++ {
		outputValue.Index(i).Set(valueValue)
	}

	return outputValue.Interface(), nil
}
