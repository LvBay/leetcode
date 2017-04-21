package substring_test

import (
	"leetcode/substring"
	"testing"
)

func TestSubstring2(t *testing.T) {
	input := []string{"abcdd", "abccb", "abccdfg", "abccdb", "ohomm"}
	out := []int{4, 3, 4, 3, 3} {
		t.Log("i:", i)
		res := substring.Substring2(v)
		if res != out[i] {
			t.Errorf(" %s want %d, but get %d", input[i], out[i], res)
		}
	}
}
