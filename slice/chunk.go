package slice

import (
	"errors"
	"reflect"
)

func Chunk(slice interface{}, size int) (interface{}, error) {

	if reflect.ValueOf(slice).Kind() != reflect.Slice {
		return nil, errors.New("not a slice")
	}

	len := reflect.ValueOf(slice).Len()

	if len == 0 {
		return nil, errors.New("slice is empty")
	}

	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}

	// TODO: chose better name
	numberOfChunkableValues := (len / size) * size

	chunks := make([]interface{}, 0)

	for i := 0; i < numberOfChunkableValues; i += size {
		value := reflect.ValueOf(slice).Slice(i, i+size).Interface()
		chunks = append(chunks, value)
	}

	// append left over values
	if numberOfChunkableValues != len {

		leftOverChunk := make([]interface{}, 0)
		for numberOfChunkableValues < len {

			value := reflect.ValueOf(slice).Index(numberOfChunkableValues).Interface()
			leftOverChunk = append(leftOverChunk, value)

			numberOfChunkableValues++
		}

		chunks = append(chunks, leftOverChunk)

	}

	return chunks, nil

}
