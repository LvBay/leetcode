package tree

import (
	"fmt"
)

// pre-order and range
func flatten(root *TreeNode) {
	p := root
	stack := []*TreeNode{}
	tmpStack := []*TreeNode{}
	for p != nil || len(tmpStack) > 0 {
		for p != nil {
			stack = append(stack, p)
			fmt.Println("append:", p.Val)
			tmpStack = append(tmpStack, p)
			p = p.Left
		}
		if len(tmpStack) > 0 {
			tmp := tmpStack[len(tmpStack)-1]
			tmpStack = tmpStack[:len(tmpStack)-1]
			p = tmp.Right
		}
	}
	p = root
	for _, v := range stack[1:] {
		p.Right = v
		p.Left = nil
		p = p.Right
	}
	ret := preorderTraversal2(root)
	fmt.Println("ret:", ret)
}
