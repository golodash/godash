package slices

import (
	"github.com/gotorn/godash/internal"
)

// Removes falsey items from slice except values you mentioned.
//
// Falsey items are {"", nil, 0, false}
func Compact(slice interface{}, excepts ...interface{}) ([]interface{}, error) {
	s, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}

	defaultFalsey := []interface{}{"", nil, 0, false}
	falsey := []interface{}{}

	for i := 0; i < len(defaultFalsey); i++ {
		remain := true
		for j := 0; j < len(excepts); j++ {
			if defaultFalsey[i] == excepts[j] {
				remain = false
			}
		}
		if remain {
			falsey = append(falsey, defaultFalsey[i])
		}
	}

	result := []interface{}{}
	j := 0
	length := len(s)

	for i := 0; i < length; i++ {
		for k := 0; k < len(falsey); k++ {
			if s[i] == falsey[k] {
				if i == j {
					j = i + 1
					continue
				}
				result = append(result, s[j:i]...)
				j = i + 1
			}
		}
	}

	if j < len(s) {
		result = append(result, s[j:length]...)
	}

	return result, nil
}
