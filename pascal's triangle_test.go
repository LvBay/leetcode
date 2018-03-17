package leetcode

import "testing"

func TestPascalriangle(t *testing.T) {
	nums := []int{5}
	for _, v := range nums {
		ret := help(v)
		t.Log(v)
		for _, v := range ret {
			t.Log(v)
		}
	}
}
