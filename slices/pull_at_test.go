package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TPullAt struct {
	name      string
	arr       interface{}
	rems      []int
	wantSlice interface{}
	wantRems  interface{}
}

var tPullAtBenchs = []TPullAt{
	{
		name: "10",
		arr:  []int{},
		rems: []int{},
	},
	{
		name: "100",
		arr:  []int{},
		rems: []int{},
	},
	{
		name: "1000",
		arr:  []int{},
		rems: []int{},
	},
	{
		name: "10000",
		arr:  []int{},
		rems: []int{},
	},
	{
		name: "100000",
		arr:  []int{},
		rems: []int{},
	},
}

func init() {
	for j := 0; j < len(tPullAtBenchs); j++ {
		length, _ := strconv.Atoi(tPullAtBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tPullAtBenchs[j].arr = append(tPullAtBenchs[j].arr.([]int), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
			tPullAtBenchs[j].rems = append(tPullAtBenchs[j].rems, 0+(i*10), 1+(i*10), 2+(i*10), 3+(i*10), 4+(i*10), 5+(i*10))
		}
	}
}

func TestPullAt(t *testing.T) {
	var tests = []TPullAt{
		{
			name:      "nil",
			arr:       nil,
			rems:      nil,
			wantSlice: nil,
			wantRems:  nil,
		},
		{
			name:      "empty",
			arr:       []int{},
			rems:      []int{},
			wantSlice: []int{},
			wantRems:  []int{},
		},
		{
			name:      "normal",
			arr:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			rems:      []int{0, 1, 2},
			wantSlice: []int{3, 4, 5, 6, 7, 8, 9},
			wantRems:  []int{0, 1, 2},
		},
		{
			name:      "shuffle",
			arr:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			rems:      []int{0, 4, 9, 4, 9, 0},
			wantSlice: []int{1, 2, 3, 5, 6, 7, 8},
			wantRems:  []int{0, 4, 9},
		},
		{
			name:      "all remove",
			arr:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			rems:      []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			wantSlice: []int{},
			wantRems:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer func(t *testing.T, wantSlice, wantRems interface{}) {
				err := recover()

				if err != nil && wantSlice != nil && wantRems != nil {
					t.Errorf("wantSlice = %v, wantRems = %v, err = %s", wantSlice, wantRems, err)
				}
			}(t, subject.wantSlice, subject.wantRems)
			gotSlice, gotRems := PullAt(subject.arr, subject.rems)

			if ok := internal.Same(gotSlice, subject.wantSlice); !ok {
				t.Errorf("gotSlice = %v, gotRem = %v, wantSlice = %v, wantRems = %v", gotSlice, gotRems, subject.wantSlice, subject.wantRems)
				return
			}

			if ok := internal.Same(gotRems, subject.wantRems); !ok {
				t.Errorf("gotSlice = %v, gotRem = %v, wantSlice = %v, wantRems = %v", gotSlice, gotRems, subject.wantSlice, subject.wantRems)
				return
			}
		})
	}
}

func BenchmarkPullAt(b *testing.B) {
	for j := 0; j < len(tPullAtBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tPullAtBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PullAt(tPullAtBenchs[j].arr, tPullAtBenchs[j].rems)
			}
		})
	}
}
