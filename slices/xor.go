package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

func Xor(slices ...interface{}) (interface{}, error) {
	var length int = 0
	for i := 0; i < len(slices); i++ {
		if err := internal.SliceCheck(slices[i]); err != nil {
			return nil, err
		}
		length += reflect.ValueOf(slices[i]).Len()
	}

	unSeenItems := make([]interface{}, 0, length)
	var sCheck bool
	for i := 0; i < len(slices); i++ {
		item := reflect.ValueOf(slices[i])
		for j := 0; j < item.Len(); j++ {
			sCheck = true
			for sIndex, s := range unSeenItems {
				if ok, _ := same(item.Index(j).Interface(), s); ok {
					unSeenItems = append(unSeenItems[:sIndex], unSeenItems[sIndex+1:]...)
					sCheck = false
					break
				}
			}
			if sCheck {
				unSeenItems = append(unSeenItems, item.Index(j).Interface())
			}
		}
	}

	return unSeenItems, nil
}
