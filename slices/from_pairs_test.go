package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TFromPairs struct {
	name string
	arr  interface{}
	want interface{}
}

var tFromPairsBenchs = []TFromPairs{
	{
		name: "10",
		arr:  [][]interface{}{},
	},
	{
		name: "100",
		arr:  [][]interface{}{},
	},
	{
		name: "1000",
		arr:  [][]interface{}{},
	},
	{
		name: "10000",
		arr:  [][]interface{}{},
	},
	{
		name: "100000",
		arr:  [][]interface{}{},
	},
}

func init() {
	for j := 0; j < len(tFromPairsBenchs); j++ {
		length, _ := strconv.Atoi(tFromPairsBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tFromPairsBenchs[j].arr = append(tFromPairsBenchs[j].arr.([][]interface{}), [][]interface{}{{"0", 1}, {"2", 3}, {"4", 5}, {"6", 7}, {"8", 9}}...)
		}
	}
}

func TestFromPairs(t *testing.T) {
	var tests = []TFromPairs{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  [][]interface{}{},
			want: map[interface{}]interface{}{},
		},
		{
			name: "none",
			arr:  [][]interface{}{{}, {}, {"4", 5, 76}, {"6", 7, "*88"}, {}},
			want: map[interface{}]interface{}{},
		},
		{
			name: "normal",
			arr:  [][]interface{}{{"0", 1}, {"2", 3}, {"4", 5}, {"6", 7}, {"8", 9}},
			want: map[interface{}]interface{}{"0": 1, "2": 3, "4": 5, "6": 7, "8": 9},
		},
		{
			name: "type based",
			arr:  [][]string{{"0", "1"}, {"2", "3"}, {"4", "5"}, {"6", "7"}, {"8", "9"}},
			want: map[string]string{"0": "1", "2": "3", "4": "5", "6": "7", "8": "9"},
		},
		{
			name: "multiple",
			arr:  [][]interface{}{{"0", 1}, {"2", 3, 45, 98, "["}, {"4"}, {"5", 6}, {"8", 9}},
			want: map[interface{}]interface{}{"0": 1, "4": nil, "5": 6, "8": 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := FromPairs(subject.arr)

			if !generals.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkFromPairs(b *testing.B) {
	for j := 0; j < len(tFromPairsBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tFromPairsBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FromPairs(tFromPairsBenchs[j].arr)
			}
		})
	}
}
