package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

func Join(slice []string, separator string) (string, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return "", err
	}
	if reflect.TypeOf(separator).Kind() != reflect.String {
		return "", errors.New("separator should be a string")
	}
	for i := 0; i < len(slice); i++ {
		if reflect.TypeOf(slice[i]).Kind() != reflect.String {
			return "", errors.New("slices should be strings")
		}
		continue
	}
	if len(slice) == 0 {
		return "", nil
	}
	var result string
	for i := 0; i < len(slice); i++ {
		if len(result) == 0 {
			result += slice[0]
		} else {
			result += separator + slice[i]
		}
	}
	return result, nil
}
