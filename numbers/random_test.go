package maths

import (
	"testing"

	"github.com/golodash/godash/internal"
)

type TRandom struct {
	name     string
	lower    interface{}
	upper    interface{}
	floating bool
	want     interface{}
}

var tRandomBench = TRandom{
	lower: 1355,
	upper: 4325,
}

func TestRandom(t *testing.T) {
	var tests = []TRandom{
		{
			name:     "empty",
			lower:    0,
			upper:    0,
			floating: false,
			want:     int64(0),
		},
		{
			name:     "int int int",
			lower:    0,
			upper:    1,
			floating: false,
			want:     int64(0),
		},
		{
			name:     "int int int non zero",
			lower:    1,
			upper:    1,
			floating: false,
			want:     int64(1),
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Random(subject.lower, subject.upper, subject.floating)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Random(tRandomBench.lower, tRandomBench.upper, tRandomBench.floating)
	}
}
