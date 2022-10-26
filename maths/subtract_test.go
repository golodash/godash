package maths

import (
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TSubtract struct {
	name    string
	number1 interface{}
	number2 interface{}
	want    interface{}
}

var tSubtractBench = TSubtract{
	number1: 0.00026546597,
	number2: 4567,
}

func TestSubtract(t *testing.T) {
	var tests = []TSubtract{
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
			number1: 15,
			number2: 3,
			want:    12,
		},
		{
			name:    "int float64",
			number1: 3,
			number2: 1.231,
			want:    1.769,
		},
		{
			name:    "float64 float64",
			number1: 15.1,
			number2: 3.231,
			want:    11.869,
		},
		{
			name:    "uint8 uint16",
			number1: uint8(15),
			number2: uint16(3),
			want:    uint16(12),
		},
		{
			name:    "int8 uint16",
			number1: int8(15),
			number2: uint16(3),
			want:    uint16(12),
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Subtract(subject.number1, subject.number2)

			if !generals.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(tSubtractBench.number1, tSubtractBench.number2)
	}
}
