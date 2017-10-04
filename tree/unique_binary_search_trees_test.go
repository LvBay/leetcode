package main

import (
	"reflect"
	"testing"
)

func TestNumTrees(t *testing.T) {
	input := []int{3, 4}
	except := []int{5, 14}

	for i, v := range input {
		if ret := numTrees(v); ret != except[i] {
			t.Errorf("except:%v,but get %v", except[i], ret)
		}
	}
}

func TestIsValidBST(t *testing.T) {
	input := []*TreeNode{
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
	}
	except := []bool{false, true}

	for i, v := range input {
		if ret := isValidBST(v); ret != except[i] {
			t.Errorf("except:%v,but get %v", except[i], ret)
		}
	}
}

func TestRecoverTree(t *testing.T) {
	input := []*TreeNode{
		&TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
	}
	except := []*TreeNode{
		&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
	}

	for i, v := range input {
		if recoverTree(v); !reflect.DeepEqual(v, except[i]) {
			t.Errorf("except:%+v,but get %+v", except[i], v)
		}
	}
}

func TestLevelOrder(t *testing.T) {
	input := []*TreeNode{
		&TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
	}
	except := [][][]int{
		[][]int{[]int{2}, []int{3, 1}},
	}

	for i, v := range input {
		if ret := levelOrder(v); !reflect.DeepEqual(ret, except[i]) {
			t.Errorf("except:%+v,but get %+v", except[i], ret)
		}
	}
}

func TestZigzagLevelOrder(t *testing.T) {
	input := []*TreeNode{
		&TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 10}}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
	}
	except := [][][]int{
		[][]int{[]int{3}, []int{20, 9}, []int{2, 10, 15, 7}},
	}

	for i, v := range input {
		if ret := zigzagLevelOrder(v); !reflect.DeepEqual(ret, except[i]) {
			t.Errorf("except:%+v,but get %+v", except[i], ret)
		}
	}
}

func Test_buildTree2(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{args: args{preorder: []int{3, 9, 2, 10, 20, 15, 7}, inorder: []int{2, 9, 10, 3, 15, 20, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree2(tt.args.preorder, tt.args.inorder)
			got2 := inorderTraversal(got)
			if !reflect.DeepEqual(tt.args.inorder, got2) {
				t.Errorf("except:%+v,but get %+v", tt.args.inorder, got2)
			}
		})
	}
}

func Test_buildTree3(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{args: args{inorder: []int{2, 9, 10, 3, 15, 20, 7}, postorder: []int{2, 10, 9, 15, 7, 20, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree3(tt.args.inorder, tt.args.postorder)
			got2 := mysuorderTraversal(got)
			if !reflect.DeepEqual(tt.args.postorder, got2) {
				t.Errorf("except:%+v,but get %+v", tt.args.inorder, got2)
			}
		})
	}
}

func TestLevelOrderBottom(t *testing.T) {
	input := []*TreeNode{
		&TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 10}}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
	}
	except := [][][]int{
		[][]int{[]int{2, 10, 15, 7}, []int{9, 20}, []int{3}},
		[][]int{[]int{2}, []int{1}},
	}

	for i, v := range input {
		if ret := levelOrderBottom(v); !reflect.DeepEqual(ret, except[i]) {
			t.Errorf("except:%+v,but get %+v", except[i], ret)
		}
	}
}

func TestSortedArrayToBST(t *testing.T) {
	input := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2},
		[]int{1, 2, 3},
	}

	for _, v := range input {
		if ret := sortedArrayToBST(v); !reflect.DeepEqual(inorderTraversal(ret), v) {
			t.Errorf("except:%+v,but get %+v", v, inorderTraversal(ret))
		}
	}
}

func TestIsBalanced1(t *testing.T) {
	input := []*TreeNode{
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
	}

	except := []bool{
		false,
	}

	for i, v := range input {
		if ret := isBalanced1(v); ret != except[i] {
			t.Errorf("except:%+v,but get %+v", except[i], ret)
		}
	}
}

func TestIsBalanced2(t *testing.T) {
	input := []*TreeNode{
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
	}

	except := []bool{
		false,
	}

	for i, v := range input {
		if ret := isBalanced2(v); ret != except[i] {
			t.Errorf("except:%+v,but get %+v", except[i], ret)
		}
	}
}

func TestMinDepth(t *testing.T) {
	input := []*TreeNode{
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
	}

	except := []int{
		3,
		2,
	}

	for i, v := range input {
		if ret := minDepth(v); ret != except[i] {
			t.Errorf("except:%+v,but get %+v", except[i], ret)
		}
	}
}
