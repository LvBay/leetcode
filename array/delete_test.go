package array

import (
	"testing"
)

func TestDelete(t *testing.T) {
	in := [][]int{
		[]int{1},
		[]int{3, 4, 4, 4},    // 2
		[]int{3, 4, 4, 4},    // 0
		[]int{4, 6, 1, 5, 2}, // 0
		[]int{4, 6, 1, 5, 2}, // 6
		[]int{4, 6, 1, 5, 2}, // 2
		[]int{4, 6, 1, 5, 2}, // 2
	}
	in2 := []int{
		1, 4, 3, 4, 2, 5, 6,
	}
	out := [][]int{
		[]int{},
		[]int{3},
		[]int{4, 4, 4},    // 0
		[]int{6, 1, 5, 2}, // 0
		[]int{4, 6, 1, 5}, // 6
		[]int{4, 6, 1, 2}, // 2
		[]int{4, 1, 5, 2}, // 2
	}

	for i := 0; i < len(in); i++ {
		get := delete(in[i], in2[i])
		if !compare(get, out[i]) {
			t.Error(i)
			return
		}
	}
}
