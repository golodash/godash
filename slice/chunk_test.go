package slice

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	chunks, err := Chunk(slice, 3)

	if err != nil {
		t.Error(err)
	}

	if reflect.ValueOf(chunks).Kind() != reflect.Slice {
		t.Error("not a slice")
	}

	if reflect.ValueOf(chunks).Len() != 4 {
		t.Error("not a slice")
	}

}
