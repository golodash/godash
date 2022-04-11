package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TLatest struct {
	name string
	arr  []interface{}
	want interface{}
}

var tLatestBenchs = []TLatest{
	{
		name: "10",
		arr:  []interface{}{},
	},
	{
		name: "100",
		arr:  []interface{}{},
	},
	{
		name: "1000",
		arr:  []interface{}{},
	},
	{
		name: "10000",
		arr:  []interface{}{},
	},
	{
		name: "100000",
		arr:  []interface{}{},
	},
	{
		name: "1000000",
		arr:  []interface{}{},
	},
}

func init() {
	for j := 0; j < len(tLatestBenchs); j++ {
		length, _ := strconv.Atoi(tLatestBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tLatestBenchs[j].arr = append(tLatestBenchs[j].arr, []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestLatest(t *testing.T) {
	tests := []TLatest{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []interface{}{},
			want: nil,
		},
		{
			name: "normal",
			arr:  []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 9,
		},
	}

	for _, subject := range tests {
		got, err := Latest(subject.arr)
		if err != nil {
			if subject.want != nil {
				t.Errorf("Latest() got = %v, wanted = %v", got, subject.want)
			}
			return
		}

		if got != subject.want {
			t.Errorf("Latest() got = %v, wanted = %v", got, subject.want)
			return
		}
	}
}

func BenchmarkLatest(b *testing.B) {
	for j := 0; j < len(tLatestBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tLatestBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Latest(tLatestBenchs[j].arr)
			}
		})
	}
}
