package math

import (
	"fmt"
	"math"
	"testing"

	"github.com/golodash/godash/internal"
)

type TPower struct {
	name   string
	input  interface{}
	number int
	want   interface{}
}

var tPowerBenchs = []TPower{
	{
		name:   "10",
		input:  0.213465,
		number: 10,
	},
	{
		name:   "100",
		input:  0.213465,
		number: 100,
	},
	{
		name:   "1000",
		input:  0.213465,
		number: 1000,
	},
	{
		name:   "10000",
		input:  0.213465,
		number: 10000,
	},
	{
		name:   "100000",
		input:  0.213465,
		number: 100000,
	},
}

func TestPower(t *testing.T) {
	var tests = []TPower{
		{
			name:   "nil",
			input:  nil,
			number: 0,
			want:   nil,
		},
		{
			name:   "empty",
			input:  0,
			number: 0,
			want:   1,
		},
		{
			name:   "positive int input, positive number",
			input:  5,
			number: 3,
			want:   125,
		},
		{
			name:   "positive int input, negative number",
			input:  5,
			number: -3,
			want:   1.0 / 125.0,
		},
		{
			name:   "positive float input, positive number",
			input:  1.1,
			number: 3,
			want:   math.Pow(1.1, 3),
		},
		{
			name:   "positive float input, negative number",
			input:  1.1,
			number: -3,
			want:   math.Pow(1.1, -3),
		},
		{
			name:   "negative int input, positive number",
			input:  -5,
			number: 3,
			want:   -125,
		},
		{
			name:   "negative int input, negative number",
			input:  -5,
			number: -3,
			want:   1.0 / -125.0,
		},
		{
			name:   "negative float input, positive number",
			input:  -1.1,
			number: 3,
			want:   math.Pow(-1.1, 3),
		},
		{
			name:   "negative float input, negative number",
			input:  -1.1,
			number: -3,
			want:   math.Pow(-1.1, -3),
		},
		{
			name:   "negative int input, positive even number",
			input:  -2,
			number: 4,
			want:   16,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Power(subject.input, subject.number)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkPower(b *testing.B) {
	for j := 0; j < len(tPowerBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tPowerBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Power(tPowerBenchs[j].input, tPowerBenchs[j].number)
			}
		})
	}
}
