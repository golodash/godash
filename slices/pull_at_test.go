package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TPullAt struct {
	name      string
	arr       []int
	rems      []int
	wantSlice []int
	wantRems  []int
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
	{
		name: "1000000",
		arr:  []int{},
		rems: []int{},
	},
}

func init() {
	for j := 0; j < len(tPullAtBenchs); j++ {
		length, _ := strconv.Atoi(tPullAtBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tPullAtBenchs[j].arr = append(tPullAtBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
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
			gotSlice, gotRems, err := PullAt(subject.arr, subject.rems)

			if err != nil {
				if subject.wantSlice != nil || subject.wantRems != nil {
					t.Errorf("PullAt() gotSlice = %v, gotRems = %v, wantSlice = %v, wantRems = %v", gotSlice, gotRems, subject.wantSlice, subject.wantRems)
				}
				return
			}

			if len(gotSlice) != len(subject.wantSlice) || len(gotRems) != len(subject.wantRems) {
				t.Errorf("PullAt() gotSlice = %v, gotRems = %v, wantSlice = %v, wantRems = %v", gotSlice, gotRems, subject.wantSlice, subject.wantRems)
				return
			}

			for i := 0; i < len(gotSlice); i++ {
				if gotSlice[i].(int) != subject.wantSlice[i] {
					t.Errorf("PullAt() gotSlice = %v, gotRems = %v, wantSlice = %v, wantRems = %v", gotSlice, gotRems, subject.wantSlice, subject.wantRems)
					return
				}
			}

			for i := 0; i < len(gotRems); i++ {
				if gotRems[i].(int) != subject.wantRems[i] {
					t.Errorf("PullAt() gotSlice = %v, gotRems = %v, wantSlice = %v, wantRems = %v", gotSlice, gotRems, subject.wantSlice, subject.wantRems)
					return
				}
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
