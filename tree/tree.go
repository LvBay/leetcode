package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret1, ret2, ret3 []int

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	p := root
	ret := []int{}
	stack := []*TreeNode{}
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = tmp.Right
			ret = append(ret, tmp.Val)
		}
	}
	return ret
}

// 中序遍历
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return ret1
	}
	inorderTraversal2(root.Left)
	ret1 = append(ret1, root.Val)
	inorderTraversal2(root.Right)
	return ret1
}

// 先序遍历
func preorderTraversal(root *TreeNode) []int {
	p := root
	ret := []int{}
	stack := []*TreeNode{}
	for p != nil || len(stack) != 0 {
		for p != nil {
			ret = append(ret, p.Val)
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = tmp.Right

		}
	}
	return ret
}

// 先序遍历
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return ret2
	}
	ret2 = append(ret2, root.Val)
	preorderTraversal2(root.Left)
	preorderTraversal2(root.Right)
	return ret2
}

// 后序遍历
func suorderTraversal(root *TreeNode) []int {
	p := root
	ret := []int{}
	stack := []*TreeNode{}
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = tmp.Right
			ret = append(ret, tmp.Val)
		}
	}
	return ret
}

// 后序遍历
func suorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return ret3
	}
	suorderTraversal2(root.Left)
	suorderTraversal2(root.Right)
	ret3 = append(ret3, root.Val)
	return ret3
}

// 构造二叉树
func buildTree(nums []int) *TreeNode {
	var ret *TreeNode
	for _, v := range nums {
		if ret == nil {
			ret = &TreeNode{Val: v}
			continue
		}
		p := ret
		for p != nil {
			if v > p.Val {
				if p.Right == nil {
					p.Right = &TreeNode{Val: v}
					break
				} else {
					p = p.Right
				}
			} else {
				if p.Left == nil {
					p.Left = &TreeNode{Val: v}
					break
				} else {
					p = p.Left
				}
			}
		}
	}
	return ret
}
