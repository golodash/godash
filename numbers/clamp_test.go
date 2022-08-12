package maths

import (
	"testing"

	"github.com/golodash/godash/internal"
)

type TClamp struct {
	name   string
	number interface{}
	lower  interface{}
	upper  interface{}
	want   interface{}
}

var tClampBench = TClamp{
	number: 243,
	lower:  1355,
	upper:  4325,
}

func TestClamp(t *testing.T) {
	var tests = []TClamp{
		{
			name:   "nil",
			number: nil,
			lower:  nil,
			upper:  nil,
			want:   nil,
		},
		{
			name:   "empty",
			number: 0,
			lower:  0,
			upper:  0,
			want:   0,
		},
		{
			name:   "int int int",
			number: -7,
			lower:  -5,
			upper:  -2,
			want:   -5,
		},
		{
			name:   "float64 float64 float64",
			number: 0.3,
			lower:  1.1,
			upper:  1.3,
			want:   1.1,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Clamp(subject.number, subject.lower, subject.upper)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkClamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Clamp(tClampBench.number, tClampBench.lower, tClampBench.upper)
	}
}
