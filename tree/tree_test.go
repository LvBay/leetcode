package main

import (
	"testing"
)

func compareSlice(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func TestBuildTree(t *testing.T) {
	nums := []int{9, 5, 13, 2, 6, 10, 15}

	root := buildTree(nums)
	if root.Val != 9 {
		t.Error("1")
	}
	if root.Left.Val != 5 {
		t.Error("2")
	}
	if root.Right.Val != 13 {
		t.Error("3")
	}

	if root.Left.Left.Val != 2 {
		t.Error("4")
	}
	if root.Left.Right.Val != 6 {
		t.Error("4")
	}
	if root.Right.Left.Val != 10 {
		t.Error("5")
	}
	if root.Right.Right.Val != 15 {
		t.Error("6")
	}

}

func TestBuildNomalTree(t *testing.T) {
	nums := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, 5, 1}

	root := buildNomalTree(nums)
	// ret := inorderTraversal(root)
	// t.Log(ret)
	if root.Val != 5 {
		t.Error("1")
	}
	if root.Left.Val != 4 {
		t.Error("2")
	}
	if root.Right.Val != 8 {
		t.Error("3")
	}

	if root.Left.Left.Val != 11 {
		t.Error("4")
	}
	if root.Left.Right != nil {
		t.Error("4")
	}
	if root.Right.Left.Val != 13 {
		t.Error("5")
	}
	if root.Right.Right.Val != 4 {
		t.Error("6")
	}

}

func TestInorder(t *testing.T) {
	nums := []int{9, 5, 13, 2, 6, 10, 15}

	root := buildTree(nums)
	x, y := inorderTraversal(root), inorderTraversal2(root)
	if !compareSlice(x, y) {
		t.Errorf("expect:%v,get:%v", y, x)
	}
}

func TestPreorder(t *testing.T) {
	nums := []int{9, 5, 13, 2, 6, 10, 15}

	root := buildTree(nums)

	x, y := preorderTraversal(root), preorderTraversal2(root)
	if !compareSlice(x, y) {
		t.Errorf("expect:%v,get:%v", y, x)
	}
}

func TestSuorder(t *testing.T) {
	nums := []int{9}

	root := buildTree(nums)

	x, y := suorderTraversal(root), suorderTraversal2(root)
	if !compareSlice(x, y) {
		t.Errorf("expect:%v,get:%v", y, x)
	}
}
