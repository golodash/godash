package maths

import (
	"testing"

	"github.com/golodash/godash/internal"
)

type TDivide struct {
	name    string
	number1 interface{}
	number2 interface{}
	want    interface{}
}

var tDivideBench = TDivide{
	number1: 0.00026546597,
	number2: 4567,
}

func TestDivide(t *testing.T) {
	var tests = []TDivide{
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
			want:    nil,
		},
		{
			name:    "int int",
			number1: 15,
			number2: 3,
			want:    5,
		},
		{
			name:    "int float64",
			number1: 15,
			number2: 1.5,
			want:    10.0,
		},
		{
			name:    "float64 float64",
			number1: 4.5,
			number2: 1.5,
			want:    3.0,
		},
		{
			name:    "uint8 uint16",
			number1: uint8(15),
			number2: uint16(3),
			want:    uint16(5),
		},
		{
			name:    "int8 uint16",
			number1: int8(15),
			number2: uint16(3),
			want:    uint16(5),
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Divide(subject.number1, subject.number2)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(tDivideBench.number1, tDivideBench.number2)
	}
}
