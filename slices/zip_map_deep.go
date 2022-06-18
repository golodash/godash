package slices

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/golodash/godash/internal"
)

// This method is like ZipMap except that it supports property paths.
//
// Complexity: O(n)
func ZipMapDeep(keys []string, values interface{}) (interface{}, error) {
	if ok := internal.SliceCheck(values); !ok {
		panic("passed 'values' variable is not slice type")
	}

	valuesValue := reflect.ValueOf(values)
	if len(keys) != valuesValue.Len() || len(keys) == 0 || valuesValue.Len() == 0 {
		return nil, errors.New("length of both 'keys' and 'values' slices has to be same and not empty")
	}

	keysValue := reflect.ValueOf(keys)
	outType, err := GetPropertyPathType(keysValue.Index(0).String(), valuesValue.Type().Elem())
	if err != nil {
		return nil, err
	}
	var output interface{}
	if outType.Kind() == reflect.Map {
		output = reflect.MakeMap(outType).Interface()
	} else if outType.Kind() == reflect.Ptr {
		value := reflect.New(outType.Elem())
		value.Elem().Set(reflect.MakeSlice(outType.Elem(), 0, 0))
		output = value.Interface()
	}
	for i := 0; i < keysValue.Len(); i++ {
		propertyKey := keysValue.Index(i).String()
		tempValue := reflect.ValueOf(output)
		if outType.Kind() == reflect.Ptr {
			tempValue = reflect.ValueOf(output).Elem()
		}
		for propertyKey != "" {
			outSlice, outKey, nextSlice, nextMap, remaining, err := nextProperty(propertyKey)
			propertyKey = remaining
			if err != nil {
				return nil, errors.New(err.Error() + fmt.Sprintf(". index = %d", i))
			} else if outSlice != -1 {
				if ok := internal.SliceCheck(tempValue.Interface()); !ok {
					panic(fmt.Sprintf("key formats are wrong. index = %d", i))
				}

				if tempValue.Len() <= outSlice {
					for j := tempValue.Len() - 1; j < outSlice; j++ {
						if outType, err := GetPropertyPathType(propertyKey, valuesValue.Type().Elem()); err != nil {
							return nil, err
						} else {
							if outType.Kind() == reflect.Ptr {
								value := reflect.New(outType.Elem())
								value.Elem().Set(reflect.MakeSlice(outType.Elem(), 0, 0))
								tempValue.Set(reflect.Append(tempValue, value))
							} else if outType.Kind() == reflect.Map {
								value := reflect.MakeMap(outType)
								tempValue.Set(reflect.Append(tempValue, value))
							} else {
								value := reflect.Zero(outType)
								tempValue.Set(reflect.Append(tempValue, value))
							}
						}
					}
				}

				if nextSlice {
					if outType, err := GetPropertyPathType(propertyKey, valuesValue.Type().Elem()); err != nil {
						return nil, err
					} else {
						if tempValue.Index(outSlice).Type() != outType {
							return nil, errors.New("key formats are wrong" + fmt.Sprintf(". index = %d", i))
						}
						if tempValue.Index(outSlice).Interface() == nil {
							value := reflect.New(outType.Elem())
							value.Set(reflect.MakeSlice(outType.Elem(), 0, 0))
							tempValue.Index(outSlice).Set(value)
							tempValue = value.Elem()
						} else {
							tempValue = tempValue.Index(outSlice).Elem()
						}
					}
				} else if nextMap {
					if outType, err := GetPropertyPathType(propertyKey, valuesValue.Type().Elem()); err != nil {
						return nil, err
					} else {
						if tempValue.Index(outSlice).Type() != outType {
							return nil, errors.New("key formats are wrong" + fmt.Sprintf(". index = %d", i))
						}
						if tempValue.Index(outSlice).Interface() == nil {
							value := reflect.MakeMap(outType)
							tempValue.Index(outSlice).Set(value)
							tempValue = value
						} else {
							tempValue = tempValue.Index(outSlice)
						}
					}
				} else if !nextSlice && !nextMap && propertyKey == "" {
					tempValue.Index(outSlice).Set(valuesValue.Index(i))
				} else {
					return nil, errors.New("key formats are wrong" + fmt.Sprintf(". index = %d", i))
				}
			} else if outKey != "" {
				if nextSlice {
					if outType, err := GetPropertyPathType(propertyKey, valuesValue.Type().Elem()); err != nil {
						return nil, err
					} else {
						mapIndexType := tempValue.Type().Elem()
						mapIndex := tempValue.MapIndex(reflect.ValueOf(outKey))
						if mapIndexType != outType {
							return nil, errors.New("key formats are wrong" + fmt.Sprintf(". index = %d", i))
						}
						if !mapIndex.IsValid() || mapIndex.Interface() == nil {
							value := reflect.New(outType.Elem())
							value.Elem().Set(reflect.MakeSlice(outType.Elem(), 0, 0))
							tempValue.SetMapIndex(reflect.ValueOf(outKey), value)
							tempValue = value.Elem()
						} else {
							tempValue = tempValue.MapIndex(reflect.ValueOf(outKey)).Elem()
						}
					}
				} else if nextMap {
					if outType, err := GetPropertyPathType(propertyKey, valuesValue.Type().Elem()); err != nil {
						return nil, err
					} else {
						mapIndexType := tempValue.Type().Elem()
						mapIndex := tempValue.MapIndex(reflect.ValueOf(outKey))
						if mapIndexType != outType {
							return nil, errors.New("key formats are wrong" + fmt.Sprintf(". index = %d", i))
						}
						if !mapIndex.IsValid() || mapIndex.Interface() == nil {
							value := reflect.MakeMap(outType)
							tempValue.SetMapIndex(reflect.ValueOf(outKey), value)
							tempValue = value
						} else {
							tempValue = tempValue.MapIndex(reflect.ValueOf(outKey))
						}
					}
				} else if !nextSlice && !nextMap && propertyKey == "" {
					tempValue.SetMapIndex(reflect.ValueOf(outKey), valuesValue.Index(i))
				} else {
					return nil, errors.New("key formats are wrong" + fmt.Sprintf(". index = %d", i))
				}
			}
		}
	}

	return output, nil
}

func GetPropertyPathType(key string, elementType reflect.Type) (reflect.Type, error) {
	if key == "" {
		return elementType, nil
	}
	if string(key[len(key)-1]) != "~" {
		key = key + "~"
	}

	isSlice := false
	isKey := false
	i := 0
	for i = 0; i < len(key); i++ {
		single := string(key[i])
		if isSlice {
			if single == "]" || single == "~" {
				if t, err := GetPropertyPathType(key[i+1:], elementType); err != nil {
					return nil, err
				} else {
					return reflect.PtrTo(reflect.SliceOf(t)), nil
				}
			} else if single == "." || single == "[" {
				return nil, errors.New("key formats are wrong" + fmt.Sprintf(". index = %d", i))
			} else {
				continue
			}
		} else if isKey {
			if single == "." || single == "[" || single == "~" {
				if t, err := GetPropertyPathType(key[i:], elementType); err != nil {
					return nil, err
				} else {
					return reflect.MapOf(reflect.TypeOf(""), t), nil
				}
			} else if single == "]" {
				return nil, errors.New("key formats are wrong" + fmt.Sprintf(". index = %d", i))
			} else if i == len(key)-1 {
				if t, err := GetPropertyPathType(key[i+1:], elementType); err != nil {
					return nil, err
				} else {
					return reflect.MapOf(reflect.TypeOf(""), t), nil
				}
			} else {
				continue
			}
		} else {
			if single == "[" {
				isSlice = true
			} else if single == "." {
				isKey = true
			} else if single == "~" {
				break
			} else {
				isKey = true
			}
		}
	}

	return elementType, nil
}

func nextProperty(key string) (outSlice int, outKey string, nextSlice, nextMap bool, remaining string, err error) {
	if string(key[len(key)-1]) != "~" {
		key = key + "~"
	}
	removeTilde := func(key string) string {
		if string(key[len(key)-1]) == "~" {
			return key[:len(key)-1]
		}
		return key
	}

	outSlice = -1
	outKey = ""
	remaining = ""
	nextSlice = false
	nextMap = false
	err = nil

	from := -1
	isSlice := false
	isKey := false
	for i := 0; i < len(key); i++ {
		single := string(key[i])
		if isSlice {
			if single == "]" {
				if from == i {
					err = errors.New("key formats are wrong")
					return
				}
				outSlice, err = strconv.Atoi(key[from:i])
				remaining = removeTilde(key)[i+1:]
				if err != nil {
					err = errors.New("key formats are wrong")
					return
				}
				if i+1 < len(key) {
					tempSingle := string(key[i+1])
					if tempSingle == "[" {
						nextSlice = true
					} else if tempSingle == "." {
						nextMap = true
					}
				}
				return
			} else if single == "." || single == "[" || single == "~" {
				err = errors.New("key formats are wrong")
				return
			} else {
				continue
			}
		} else if isKey {
			if single == "." || single == "[" || single == "~" {
				if from == i {
					err = errors.New("key formats are wrong")
					return
				}
				outKey = key[from:i]
				remaining = removeTilde(key)[i:]
				if single == "[" {
					nextSlice = true
				} else if single == "." {
					nextMap = true
				}
				return
			} else if single == "]" {
				err = errors.New("key formats are wrong")
				return
			} else if i == len(key)-1 {
				outKey = key[from:]
				return
			} else {
				continue
			}
		} else {
			if single == "[" {
				isSlice = true
				from = i + 1
			} else if single == "." {
				isKey = true
				from = i + 1
			} else if single == "~" {
				return
			} else {
				isKey = true
				from = i
			}
		}
	}

	err = errors.New("key formats are wrong")
	return
}
