package slice

import "testing"

type table struct {
	data   interface{}
	latest interface{}
}

func BenchmarkTestLatest(t *testing.B) {

	for i := 0; i < t.N; i++ {
		Latest([]interface{}{1, 2, 3})
	}

}

func TestLatest(t *testing.T) {

	tabels := []table{

		{
			data:   []int{1, 2, 3},
			latest: 3,
		},
		{
			data:   []string{"one", "two"},
			latest: "two",
		},
		{
			data:   []bool{true, false},
			latest: false,
		},
	}

	for _, table := range tabels {

		result, err := Latest(table.data)

		if err != nil {
			t.Error("cannot handel", table.data)
		}

		if result != table.latest {
			t.Errorf("expect %s but got %s", table.latest, result)
		}

	}

}
