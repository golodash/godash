package functions

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TWrapFunc struct {
	name     string
	inputs   interface{}
	function interface{}
	expected interface{}
}

var tWrapFuncBenchs = []TWrapFunc{
	{
		name:   "10",
		inputs: []interface{}{},
	},
	{
		name:   "100",
		inputs: []interface{}{},
	},
	{
		name:   "1000",
		inputs: []interface{}{},
	},
	{
		name:   "10000",
		inputs: []interface{}{},
	},
	{
		name:   "100000",
		inputs: []interface{}{},
	},
}

func init() {
	for j := 0; j < len(tWrapFuncBenchs); j++ {
		length, _ := strconv.Atoi(tWrapFuncBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tWrapFuncBenchs[j].inputs = append(tWrapFuncBenchs[j].inputs.([]interface{}), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
	}
}

func TestWrapFunc(t *testing.T) {
	var tests = []TWrapFunc{
		{
			name:     "nil",
			inputs:   nil,
			function: nil,
			expected: nil,
		},
		{
			name:     "empty",
			inputs:   []interface{}{},
			function: func(inputs ...interface{}) {},
			expected: nil,
		},
		{
			name:     "empty-noInput",
			inputs:   []interface{}{},
			function: func() {},
			expected: true,
		},
		{
			name:     "inputsBigger-withVariadic",
			inputs:   []interface{}{1, 2, 3, 4},
			function: func(first interface{}, seconds ...interface{}) {},
			expected: true,
		},
		{
			name:     "inputsBigger-withoutVariadic",
			inputs:   []interface{}{1, 2, 3, 4},
			function: func(first, seconds interface{}) {},
			expected: nil,
		},
		{
			name:     "inFuncBigger-withVariadic-1",
			inputs:   []interface{}{1, 2},
			function: func(first, seconds interface{}, others ...interface{}) {},
			expected: true,
		},
		{
			name:     "inFuncBigger-withVariadic-2",
			inputs:   []interface{}{1},
			function: func(first, seconds interface{}, others ...interface{}) {},
			expected: nil,
		},
		{
			name:     "inFuncBigger-withoutVariadic",
			inputs:   []interface{}{1},
			function: func(first, seconds interface{}) {},
			expected: nil,
		},
		{
			name:     "sameLength-withVariadic",
			inputs:   []interface{}{1, 2, 3},
			function: func(first, seconds interface{}, others ...interface{}) {},
			expected: true,
		},
		{
			name:     "sameLength-withoutVariadic",
			inputs:   []interface{}{1, 2},
			function: func(first, seconds interface{}) {},
			expected: true,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.expected)
			var got func() []interface{} = nil
			if subject.inputs == nil {
				got = WrapFunc(subject.function, nil)
			} else {
				got = WrapFunc(subject.function, subject.inputs.([]interface{})...)
			}

			if subject.expected != nil && got == nil {
				t.Errorf("got = a function, wanted = %v", subject.expected)
			} else if got != nil {
				defer func() {
					e := recover()
					if e != nil {
						panic("you are fucked")
					}
				}()
				got()
			}
		})
	}
}

func benchFunc(inputs ...interface{}) {}

func BenchmarkWrapFunc(b *testing.B) {
	for j := 0; j < len(tWrapFuncBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tWrapFuncBenchs[j].name), func(b *testing.B) {
			inputs := tWrapFuncBenchs[j].inputs.([]interface{})
			for i := 0; i < b.N; i++ {
				WrapFunc(benchFunc, inputs...)
			}
		})
	}
}
