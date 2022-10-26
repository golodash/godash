package generals

import (
	"reflect"

	"github.com/golodash/godash/internal"
	"github.com/jinzhu/copier"
)

func Duplicate(input interface{}) interface{} {
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	var secondInput reflect.Value = reflect.New(inputType)
	switch inputType.Kind() {
	case reflect.Slice:
		secondInput = reflect.MakeSlice(inputType, inputValue.Len(), inputValue.Len())
	case reflect.Map:
		secondInput = reflect.MakeMapWithSize(inputType, inputValue.Len())
	case reflect.Array:
		secondInput = reflect.New(reflect.ArrayOf(inputValue.Len(), inputType.Elem())).Elem()
	}
	if internal.IsPrimitive(inputType.Kind()) {
		second := input
		return second
	}

	second := secondInput.Interface()
	copier.Copy(&second, input)
	return second
}
