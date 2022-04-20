package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TLastIndexOf struct {
	name  string
	arr   []int
	value int
	want  int
}

var tLastIndexOfBenchs = []TLastIndexOf{
	{
		name: "10",
		arr:  []int{},
	},
	{
		name: "100",
		arr:  []int{},
	},
	{
		name: "1000",
		arr:  []int{},
	},
	{
		name: "10000",
		arr:  []int{},
	},
	{
		name: "100000",
		arr:  []int{},
	},
	{
		name: "1000000",
		arr:  []int{},
	},
}

func init() {
	for j := 0; j < len(tLastIndexOfBenchs); j++ {
		length, _ := strconv.Atoi(tLastIndexOfBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tLastIndexOfBenchs[j].arr = append(tLastIndexOfBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
		tLastIndexOfBenchs[j].arr = append(tLastIndexOfBenchs[j].arr, 10)
		tLastIndexOfBenchs[j].value = 10
	}
}

func TestLastIndexOf(t *testing.T) {
	var tests = []TLastIndexOf{
		{
			name:  "nil",
			arr:   nil,
			value: -1,
			want:  -1,
		},
		{
			name:  "empty",
			arr:   []int{},
			value: -1,
			want:  -1,
		},
		{
			name:  "normal",
			arr:   []int{0, 1, 2, 3, 4, 5, 4, 7, 8, 9},
			value: 4,
			want:  6,
		},
		{
			name:  "does not exist",
			arr:   []int{0, 1, 2, 3, 4, 6, 7, 8, 9},
			value: 10,
			want:  -1,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := LastIndexOf(subject.arr, subject.value)
			if err != nil {
				if subject.want != -1 {
					t.Errorf("LastIndexOf() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if got != subject.want {
				t.Errorf("LastIndexOf() got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkLastIndexOf(b *testing.B) {
	for j := 0; j < len(tLastIndexOfBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tLastIndexOfBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LastIndexOf(tLastIndexOfBenchs[j].arr, tLastIndexOfBenchs[j].value)
			}
		})
	}
}
