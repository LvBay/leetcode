package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret1, ret2, ret3 []int

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	p := root
	stack := []*TreeNode{}
	ret := []int{}

	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			p = tmp.Right // 每次添加值之后，检查右子节点
			ret = append(ret, tmp.Val)
			stack = stack[:len(stack)-1]
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
	stack := []*TreeNode{}
	var cur *TreeNode
	var pre *TreeNode
	stack = append(stack, root)
	var ret []int

	for len(stack) > 0 {
		cur = stack[len(stack)-1]
		// 如果栈顶是叶子节点或者上个节点等于当前节点的子节点
		if (cur.Left == nil && cur.Right == nil) || (pre != nil && (pre == cur.Left || pre == cur.Right)) {
			ret = append(ret, cur.Val)
			stack = stack[:len(stack)-1]
			pre = cur

		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}
	return ret
}

// 后序遍历
func mysuorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	p := root
	ret := []int{}
	stack = append(stack, p)

	flag := false
	var right int
	if root.Right != nil {
		flag = true
		right = root.Right.Val
	}

	for len(stack) > 0 {
		for p != nil {
			if p.Right != nil {
				stack = append(stack, p.Right)
			}
			if p.Left != nil {
				stack = append(stack, p.Left)
			}

			p = p.Left
		}

		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			if tmp.Val == right && flag {
				p = tmp
				flag = false
				continue
			}
			stack = stack[:len(stack)-1]
			ret = append(ret, tmp.Val)

			p = tmp.Right
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

// 构造二叉树 非搜索二叉树
func buildNomalTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	var ret, node *TreeNode
	ret = &TreeNode{Val: nums[0]}
	stack := []*TreeNode{ret}
	left := true
	nextNode := true

	for _, v := range nums[1:] {
		// fmt.Println("lalal", v)
		if nextNode {
			node = stack[0]
			stack = stack[1:]
			nextNode = false
		}
		if left {
			if v == -1 {
				node.Left = nil
				// fmt.Printf("%d.Left=nil\n", node.Val)
			} else {
				node.Left = &TreeNode{Val: v}
				// fmt.Printf("%d.Left=%d\n", node.Val, v)
				stack = append(stack, node.Left)
			}

			left = false
		} else {
			if v == -1 {
				node.Right = nil
				// fmt.Printf("%d.Right=nil\n", node.Val)
			} else {
				node.Right = &TreeNode{Val: v}
				// fmt.Printf("%d.Right=%d\n", node.Val, v)
				stack = append(stack, node.Right)
			}
			nextNode = true
			left = true
		}
	}
	return ret
}
