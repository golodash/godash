package maths

import (
	"testing"

	"github.com/golodash/godash/internal"
)

type TFloor struct {
	name      string
	input     interface{}
	precision int
	want      interface{}
}

var tFloorBench = TFloor{
	input:     0.00026546597,
	precision: 5,
}

func TestFloor(t *testing.T) {
	var tests = []TFloor{
		{
			name:      "nil",
			input:     nil,
			precision: 5,
			want:      nil,
		},
		{
			name:      "empty",
			input:     0,
			precision: 0,
			want:      0,
		},
		{
			name:      "positive precision",
			input:     0.123456,
			precision: 3,
			want:      0.123,
		},
		{
			name:      "negative precision",
			input:     1532,
			precision: -3,
			want:      1000,
		},
		{
			name:      "zero precision",
			input:     1532.321,
			precision: 0,
			want:      1532.0,
		},
		{
			name:      "positive precision, negative value",
			input:     -0.123456,
			precision: 3,
			want:      -0.123,
		},
		{
			name:      "negative precision, negative value",
			input:     -1532,
			precision: -3,
			want:      -1000,
		},
		{
			name:      "zero precision, negative value",
			input:     -1532.321,
			precision: 0,
			want:      -1532.0,
		},
		{
			name:      "normal-float32 type",
			input:     float32(-1532.321),
			precision: 2,
			want:      float32(-1532.32),
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Floor(subject.input, subject.precision)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkFloor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Floor(tFloorBench.input, tFloorBench.precision)
	}
}
