package slices

import (
	"github.com/golodash/godash/internal"
)

func FromPairs(slice interface{}) (map[string]interface{}, error) {
	s, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}

	var output = map[string]interface{}{}
	for i := 0; i < len(s); i++ {
		item, err := internal.InterfaceToSlice(s[i])
		ok := false
		if err != nil {
			item, ok = s[i].([]interface{})
			if !ok {
				continue
			}
		}

		if len(item) == 2 {
			if key, ok := item[0].(string); ok {
				output[key] = item[1]
			}
		} else if len(item) == 1 {
			if key, ok := item[0].(string); ok {
				output[key] = nil
			}
		}
	}

	return output, nil
}
