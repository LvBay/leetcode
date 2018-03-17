package sort

import (
	"testing"
)

func TestQuick(t *testing.T) {
	in := [][]int{
		[]int{3, 1, 4},
		[]int{5, 4, 1, 3, 2, 1, 6, 7},
		[]int{1, 2},
		[]int{2, 1},
		[]int{1},
	}
	out := [][]int{
		[]int{1, 3, 4},
		[]int{1, 1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2},
		[]int{1, 2},
		[]int{1},
	}
	for i := 0; i < len(in); i++ {
		quickSort(in[i])
	}
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			if in[i][j] != out[i][j] {
				t.Error(in[i], out[i])
			}
		}
	}
}
