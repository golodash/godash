package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
)

type TRemoveBy struct {
	name      string
	arg1      interface{}
	wantSlice interface{}
	wantRems  interface{}
}

var TRemoveByBenchs = []TRemoveBy{
	{
		name: "10",
		arg1: []int{},
	},
	{
		name: "100",
		arg1: []int{},
	},
	{
		name: "1000",
		arg1: []int{},
	},
	{
		name: "10000",
		arg1: []int{},
	},
	{
		name: "100000",
		arg1: []int{},
	},
}

func init() {
	for i := 0; i < len(TRemoveByBenchs); i++ {
		k, _ := strconv.Atoi(TRemoveByBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TRemoveByBenchs[i].arg1 = append(TRemoveByBenchs[i].arg1.([]int), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}...)
		}
	}
}

func removeByFunctionTest(n interface{}) bool {
	return n.(int)%2 != 0
}

func TestRemoveBy(t *testing.T) {
	var tests = []TRemoveBy{
		{
			name:      "nil",
			arg1:      nil,
			wantSlice: nil,
			wantRems:  nil,
		},
		{
			name:      "empty",
			arg1:      []int{},
			wantSlice: []int{},
			wantRems:  []int{},
		},
		{
			name:      "default",
			arg1:      []int{1, 2, 3, 4, 5, 6, 7, 8, 66, 44, 5, 6, 7, 99},
			wantSlice: []int{2, 4, 6, 8, 66, 44, 6},
			wantRems:  []int{99, 7, 5, 7, 5, 3, 1},
		},
		{
			name:      "default1",
			arg1:      []int{3, 4, 5, 6, 23, 34, 56, 68, 98},
			wantSlice: []int{4, 6, 34, 56, 68, 98},
			wantRems:  []int{23, 5, 3},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer func(t *testing.T, wantSlice, wantRems interface{}) {
				err := recover()

				if err != nil && wantSlice != nil && wantRems != nil {
					t.Errorf("wantSlice = %v, wantRems = %v, err = %s", wantSlice, wantRems, err)
				}
			}(t, sample.wantSlice, sample.wantRems)
			gotSlice, gotRems := RemoveBy(sample.arg1, removeByFunctionTest)

			if !generals.Same(gotSlice, sample.wantSlice) {
				t.Errorf("gotSlice = %v, gotRem = %v, wantSlice = %v, wantRems = %v", gotSlice, gotRems, sample.wantSlice, sample.wantRems)
				return
			}

			if !generals.Same(gotRems, sample.wantRems) {
				t.Errorf("gotSlice = %v, gotRem = %v, wantSlice = %v, wantRems = %v", gotSlice, gotRems, sample.wantSlice, sample.wantRems)
				return
			}
		})
	}
}

func BenchmarkRemoveBy(b *testing.B) {
	for _, sample := range TRemoveByBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RemoveBy(sample.arg1, removeByFunctionTest)
			}
		})
	}
}
