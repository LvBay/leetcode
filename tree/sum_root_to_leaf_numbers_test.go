package tree

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	root := buildNomalTree([]int{1, 2, 3})
	ret := sumNumbers(root)
	if ret != (12 + 13) {
		t.Error("ret:", ret)
	}
}

func TestSum(t *testing.T) {
	ret := Sum([]int{1, 2, 3, 4})
	t.Log(ret)
}
