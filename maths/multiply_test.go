package maths

import (
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TMultiply struct {
	name    string
	number1 interface{}
	number2 interface{}
	want    interface{}
}

var tMultiplyBench = TMultiply{
	number1: 0.00026546597,
	number2: 4567,
}

func TestMultiply(t *testing.T) {
	var tests = []TMultiply{
		{
			name:    "nil",
			number1: nil,
			number2: nil,
			want:    nil,
		},
		{
			name:    "empty",
			number1: 0,
			number2: 0,
			want:    0,
		},
		{
			name:    "int int",
			number1: 3,
			number2: 3,
			want:    9,
		},
		{
			name:    "int float64",
			number1: 15,
			number2: 3.1,
			want:    46.5,
		},
		{
			name:    "float64 float64",
			number1: 15.1,
			number2: 3.231,
			want:    48.7881,
		},
		{
			name:    "uint8 uint16",
			number1: uint8(3),
			number2: uint16(3),
			want:    uint16(9),
		},
		{
			name:    "int8 uint16",
			number1: int8(3),
			number2: uint16(3),
			want:    uint16(9),
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Multiply(subject.number1, subject.number2)

			if !generals.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(tMultiplyBench.number1, tMultiplyBench.number2)
	}
}
