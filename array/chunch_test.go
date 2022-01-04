package array

import (
	"testing"
)

type test struct {
	name string
	arr  []any
	size int
	want [][]any
}

var tests []test

func TestChunk(t *testing.T) {
	type args struct {
		err error
	}

	tests = []test{
		test{
			name: "negative",
			arr: []any{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: -8,
			want: nil,
		},
		test{
			name: "lower size",
			arr: []any{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 0,
			want: nil,
		},
		test{
			name: "lower size",
			arr: []any{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 2,
			want: [][]any{[]any{0, 1}, []any{2, 3}, []any{4, 5}, []any{6, 7}, []any{8, 9}},
		},
		test{
			name: "half size",
			arr: []any{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 5,
			want: [][]any{[]any{0, 1, 2, 3, 4}, []any{5, 6, 7, 8, 9}},
		},
		test{
			name: "more than half",
			arr: []any{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 8,
			want: [][]any{[]any{0, 1, 2, 3, 4, 5, 6, 7}, []any{8, 9}},
		},
		test{
			name: "equal to size",
			arr: []any{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 10,
			want: [][]any{[]any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		test{
			name: "more than size",
			arr: []any{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 15,
			want: [][]any{[]any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := Chunk(subject.arr, subject.size)

			if err != nil && got == nil && subject.want == nil {
				return
			} else if got == nil {
				t.Errorf("%v() got = %v, wanted = %v", "Chunk", got, subject.want)
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("%v() got = %v, wanted = %v", "Chunk", got, subject.want)
				return
			}
			for i := 0; i < len(got); i++ {
				if len(got[i]) != len(subject.want[i]) {
					t.Errorf("%v() got = %v, wanted = %v", "Chunk", got, subject.want)
					return
				}
			}

			for i := 0; i < len(got); i++ {
				for j := 0; j < len(got[i]); j++ {
					if got[i][j] != subject.want[i][j] {
						t.Errorf("%v() got = %v, wanted = %v", "Chunk", got, subject.want)
						return
					}
				}
			}
		})
	}
}

func BenchmarkChunk(b *testing.B) {

	var test = test{
		name: "",
		arr:  []any{},
		size: 2,
		want: nil,
	}

	for i := 0; i < 100; i++ {
		test.arr = append(test.arr, []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	}

	for i := 0; i < b.N; i++ {
		Chunk(test.arr, test.size)
	}
}
