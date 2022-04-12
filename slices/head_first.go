package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

func Head(slice interface{}) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	s := reflect.ValueOf(slice)

	if s.Len() == 0 {
		return nil, errors.New("slice is empty")
	}

	return s.Index(0).Interface(), nil
}

func First(slice interface{}) (interface{}, error) {
	return Head(slice)
}