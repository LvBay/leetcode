package tree

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	root := buildNomalTree([]int{1, 2, 5, 3, 4, -1, 6})
	t.Log(root.Val, root.Left.Val, root.Right.Val)
	flatten(root)
	if root.Val != 1 {
		t.Error("1")
		return
	}
	if root.Right == nil || root.Right.Val != 2 {
		t.Error("2")
		return
	}
	if root.Right.Right == nil || root.Right.Right.Val != 3 {
		t.Error("3")
		return
	}
	if root.Right.Right.Right == nil || root.Right.Right.Right.Val != 4 {
		t.Error("4")
		return
	}
	if root.Right.Right.Right.Right == nil || root.Right.Right.Right.Right.Val != 5 {
		t.Error("5")
		return
	}
	if root.Right.Right.Right.Right.Right == nil || root.Right.Right.Right.Right.Right.Val != 6 {
		t.Error("6")
		return
	}
}
