package numbers

import (
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TInRange struct {
	name   string
	number interface{}
	start  interface{}
	end    interface{}
	want   bool
}

var tInRangeBench = TInRange{
	number: 1200,
	start:  1355,
	end:    4325,
}

func TestInRange(t *testing.T) {
	var tests = []TInRange{
		{
			name:   "empty",
			number: 0,
			start:  0,
			end:    0,
			want:   true,
		},
		{
			name:   "int int int",
			number: -4,
			start:  -5,
			end:    -2,
			want:   true,
		},
		{
			name:   "int int int",
			number: -7,
			start:  -5,
			end:    -2,
			want:   false,
		},
		{
			name:   "float64 int int",
			number: 2.1,
			start:  1,
			end:    2,
			want:   false,
		},
		{
			name:   "int float64 float64",
			number: 2,
			start:  1.1,
			end:    4.6,
			want:   true,
		},
		{
			name:   "int float64 float64",
			number: 1,
			start:  1.1,
			end:    4.6,
			want:   false,
		},
		{
			name:   "float64 float64 float64",
			number: 0.3,
			start:  1.1,
			end:    1.3,
			want:   false,
		},
		{
			name:   "float64 float64 float64",
			number: 1.2,
			start:  1.1,
			end:    1.3,
			want:   true,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := InRange(subject.number, subject.start, subject.end)

			if !generals.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InRange(tInRangeBench.number, tInRangeBench.start, tInRangeBench.end)
	}
}
