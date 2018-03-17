package array

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	in := [][]int{
		[]int{},
		[]int{1},    // 2
		[]int{1},    // 0
		[]int{1, 5}, // 0
		[]int{1, 5}, // 6
		[]int{1, 5}, // 2
	}
	in2 := []int{
		1, 2, 0, 0, 6, 2,
	}
	out := [][]int{
		[]int{1},
		[]int{1, 2},
		[]int{0, 1},
		[]int{0, 1, 5},
		[]int{1, 5, 6},
		[]int{1, 2, 5},
	}

	for i := 0; i < len(in); i++ {
		get := insert(in[i], in2[i])
		if !compare(get, out[i]) {
			t.Error(i)
			return
		}
	}
}

func compare(x, y interface{}) bool {
	return reflect.DeepEqual(x, y)
}
