package array

import "testing"

func TestContainsDuplicate(t *testing.T) {
	in := [][]int{
		[]int{1, 5, -2, -4, 0},
		// []int{1},
		// []int{1, 2},
		// []int{1},
	}
	for i := 0; i < len(in); i++ {
		containsDuplicate(in[i])
	}
}
