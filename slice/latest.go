package slice

import (
	"errors"
	"reflect"
)

func Latest(slice interface{}) (interface{}, error) {

	if reflect.ValueOf(slice).Kind() != reflect.Slice {
		return nil, errors.New("not a slice")
	}

	len := reflect.ValueOf(slice).Len()

	if len == 0 {
		return nil, errors.New("slice is empty")
	}

	return reflect.ValueOf(slice).Index(len - 1).Interface(), nil

}
