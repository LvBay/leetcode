package array

import "testing"

func TestRotateImage(t *testing.T) {
	in := [][][]int{
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
	}
	for i := 0; i < len(in); i++ {
		rotateImage(in[i])
	}
}
