package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TCompact struct {
	name   string
	arr    []interface{}
	except []interface{}
	want   []interface{}
}

var tCompactBenchs = []TCompact{
	{
		name:   "10",
		arr:    []interface{}{},
		except: []interface{}{},
	},
	{
		name:   "100",
		arr:    []interface{}{},
		except: []interface{}{},
	},
	{
		name:   "1000",
		arr:    []interface{}{},
		except: []interface{}{},
	},
	{
		name:   "10000",
		arr:    []interface{}{},
		except: []interface{}{},
	},
	{
		name:   "100000",
		arr:    []interface{}{},
		except: []interface{}{},
	},
	{
		name:   "1000000",
		arr:    []interface{}{},
		except: []interface{}{},
	},
}

func init() {
	for j := 0; j < len(tCompactBenchs); j++ {
		length, _ := strconv.Atoi(tCompactBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tCompactBenchs[j].arr = append(tCompactBenchs[j].arr, []interface{}{0, nil, 2, false, 4, 5, "", nil, 8, 9}...)
		}
	}
}

func TestCompact(t *testing.T) {
	var tests = []TCompact{
		{
			name:   "nil",
			arr:    nil,
			except: []interface{}{},
			want:   nil,
		},
		{
			name:   "empty",
			arr:    []interface{}{},
			except: []interface{}{},
			want:   []interface{}{},
		},
		{
			name:   "normal",
			arr:    []interface{}{0, 1, 2, 3, nil, 5, '6', 0, false, ""},
			except: []interface{}{},
			want:   []interface{}{1, 2, 3, 5, '6'},
		},
		{
			name:   "all remove",
			arr:    []interface{}{0, nil, 0, false, nil, "", nil, 0, false, ""},
			except: []interface{}{},
			want:   []interface{}{},
		},
		{
			name:   "except on",
			arr:    []interface{}{0, nil, 0, false, nil, "", nil, 0, false, ""},
			except: []interface{}{0, ""},
			want:   []interface{}{0, 0, "", 0, ""},
		},
		{
			name:   "all except on",
			arr:    []interface{}{0, nil, 0, false, nil, "", nil, 0, false, ""},
			except: []interface{}{0, "", nil, false},
			want:   []interface{}{0, nil, 0, false, nil, "", nil, 0, false, ""},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := Compact(subject.arr, subject.except...)
			if err != nil {
				if subject.want != nil {
					t.Errorf("Compact() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("Compact() got = %v, wanted = %v", got, subject.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if got[i] != subject.want[i] {
					t.Errorf("Compact() got = %v, wanted = %v", got, subject.want)
					return
				}
			}
		})
	}
}

func BenchmarkCompact(b *testing.B) {
	for j := 0; j < len(tCompactBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tCompactBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Compact(tCompactBenchs[j].arr, tCompactBenchs[j].except...)
			}
		})
	}
}
