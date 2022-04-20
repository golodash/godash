package slices

import (
	"testing"
)

type TNth struct {
	name     string
	arg1     interface{}
	arg2     int
	expected interface{}
}

var tests = []TNth{
	{
		name:     "nil",
		arg1:     nil,
		arg2:     4,
		expected: nil,
	},
	{
		name:     "empty",
		arg1:     []string{},
		arg2:     -1,
		expected: nil,
	},
	{
		name:     "default",
		arg1:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		arg2:     5,
		expected: 6,
	},
	{
		name:     "default1",
		arg1:     []string{"a", "b", "c", "d", "e", "f", "u"},
		arg2:     -2,
		expected: "f",
	},
}

func TestNth(t *testing.T) {

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := Nth(sample.arg1, sample.arg2)
			if err != nil {
				if sample.expected != nil {
					t.Errorf("got : %v but expected : %v", got, sample.expected)
				}
				return
			}
			if got != sample.expected {
				t.Errorf("got : %v but expected : %v", got, sample.expected)
				return
			}
		})

	}

}
