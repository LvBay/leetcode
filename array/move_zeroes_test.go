package array

import "testing"

func TestMoveZeroes(t *testing.T) {
	in := [][]int{
		[]int{0, 1, 0, 3, 12},
		// []int{1, 5, -2, -4, 0},
		// []int{1},
		// []int{1, 2},
		// []int{1},
	}
	for i := 0; i < len(in); i++ {
		moveZeroes(in[i])
	}
}
