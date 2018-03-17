package leetcode

import "testing"

func TestIsUgly(t *testing.T) {
	nums := []int{-2, -14, -6, 1, 2, 3, 19, 14, -2147483648}
	for _, v := range nums {
		ret := isUgly(v)
		t.Log(v, ret)
	}
}
