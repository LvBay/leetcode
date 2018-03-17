package array

import "testing"

func TestRotate(t *testing.T) {
	in := [][]int{
		[]int{1, 2, 3},
		// []int{1, 2},
		// []int{1},
	}
	for i := 0; i < len(in); i++ {
		rotate(in[i], 1)
	}
}
