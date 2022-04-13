package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TFromPairs struct {
	name string
	arr  [][]interface{}
	want map[string]interface{}
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
	{
		name: "1000000",
		arr:  [][]interface{}{},
	},
}

func init() {
	for j := 0; j < len(tFromPairsBenchs); j++ {
		length, _ := strconv.Atoi(tFromPairsBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tFromPairsBenchs[j].arr = append(tFromPairsBenchs[j].arr, [][]interface{}{{"0", 1}, {"2", 3}, {"4", 5}, {"6", 7}, {"8", 9}}...)
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
			want: map[string]interface{}{},
		},
		{
			name: "none",
			arr:  [][]interface{}{{}, {}, {"4", 5, 76}, {"6", 7, "*88"}, {}},
			want: map[string]interface{}{},
		},
		{
			name: "normal",
			arr:  [][]interface{}{{"0", 1}, {"2", 3}, {"4", 5}, {"6", 7}, {"8", 9}},
			want: map[string]interface{}{"0": 1, "2": 3, "4": 5, "6": 7, "8": 9},
		},
		{
			name: "multiple",
			arr:  [][]interface{}{{"0", 1}, {"2", 3, 45, 98, "["}, {"4"}, {"5", 6}, {"8", 9}},
			want: map[string]interface{}{"0": 1, "4": nil, "5": 6, "8": 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := FromPairs(subject.arr)
			if err != nil {
				if subject.want != nil {
					t.Errorf("FromPairs() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("FromPairs() got = %v, wanted = %v", got, subject.want)
				return
			}

			for i := 0; i < len(subject.arr); i++ {
				for j := 0; j < 2; j++ {
					if len(subject.arr[i]) == 2 {
						if key, ok := subject.arr[i][0].(string); ok {
							res, err := same(got[key], subject.arr[i][1])
							if err != nil || !res {
								t.Errorf("FromPairs() got = %v, wanted = %v", got, subject.want)
								return
							}
						} else {
							t.Errorf("FromPairs() got = %v, wanted = %v", got, subject.want)
							return
						}
					} else if len(subject.arr[i]) == 1 {
						if key, ok := subject.arr[i][0].(string); ok {
							if got[key] != nil {
								t.Errorf("FromPairs() got = %v, wanted = %v", got, subject.want)
								return
							}
						} else {
							t.Errorf("FromPairs() got = %v, wanted = %v", got, subject.want)
							return
						}
					}
				}
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
