package math

import (
	"testing"

	"github.com/golodash/godash/internal"
)

type TRound struct {
	name      string
	input     interface{}
	precision int
	want      interface{}
}

var tRoundBench = TRound{
	input:     0.00026546597,
	precision: 5,
}

func TestRound(t *testing.T) {
	var tests = []TRound{
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
			name:      "positive precision down",
			input:     0.123456,
			precision: 3,
			want:      0.123,
		},
		{
			name:      "positive precision up",
			input:     0.123656,
			precision: 3,
			want:      0.124,
		},
		{
			name:      "negative precision down",
			input:     1432,
			precision: -3,
			want:      1000,
		},
		{
			name:      "negative precision up",
			input:     1532,
			precision: -3,
			want:      2000,
		},
		{
			name:      "zero precision down",
			input:     1532.321,
			precision: 0,
			want:      1532.0,
		},
		{
			name:      "zero precision up",
			input:     1532.921,
			precision: 0,
			want:      1533.0,
		},
		{
			name:      "positive precision, negative value down",
			input:     -0.123456,
			precision: 3,
			want:      -0.123,
		},
		{
			name:      "positive precision, negative value up",
			input:     -0.123956,
			precision: 3,
			want:      -0.124,
		},
		{
			name:      "negative precision, negative value down",
			input:     -1232,
			precision: -3,
			want:      -1000,
		},
		{
			name:      "negative precision, negative value up",
			input:     -1532,
			precision: -3,
			want:      -2000,
		},
		{
			name:      "zero precision, negative value down",
			input:     -1532.321,
			precision: 0,
			want:      -1532.0,
		},
		{
			name:      "zero precision, negative value up",
			input:     -1532.821,
			precision: 0,
			want:      -1533.0,
		},
		{
			name:      "normal-float32 type down",
			input:     float32(-1532.321),
			precision: 2,
			want:      float32(-1532.32),
		},
		{
			name:      "normal-float32 type up",
			input:     float32(-1532.326),
			precision: 2,
			want:      float32(-1532.33),
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Round(subject.input, subject.precision)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkRound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Round(tRoundBench.input, tRoundBench.precision)
	}
}
