package math

import (
	"testing"

	"github.com/golodash/godash/internal"
)

type TCeil struct {
	name      string
	input     interface{}
	precision int
	want      interface{}
}

var tCeilBench = TCeil{
	input:     0.00026546597,
	precision: 5,
}

func TestCeil(t *testing.T) {
	var tests = []TCeil{
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
			want:      1,
		},
		{
			name:      "positive precision",
			input:     0.123456,
			precision: 3,
			want:      0.124,
		},
		{
			name:      "negative precision",
			input:     1532,
			precision: -3,
			want:      2000,
		},
		{
			name:      "zero precision",
			input:     1532.321,
			precision: 0,
			want:      1533.0,
		},
		{
			name:      "positive precision, negative value",
			input:     -0.123456,
			precision: 3,
			want:      -0.124,
		},
		{
			name:      "negative precision, negative value",
			input:     -1532,
			precision: -3,
			want:      -2000,
		},
		{
			name:      "zero precision, negative value",
			input:     -1532.321,
			precision: 0,
			want:      -1533.0,
		},
		{
			name:      "normal-float32 type",
			input:     float32(-1532.321),
			precision: 2,
			want:      float32(-1532.33),
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Ceil(subject.input, subject.precision)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkCeil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ceil(tCeilBench.input, tCeilBench.precision)
	}
}
