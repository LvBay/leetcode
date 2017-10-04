package main

import (
	"sort"
)

/*
	当数组为 1，2，3，4，.. i，.. n时，基于以下原则的BST建树具有唯一性：
	以i为根节点的树，其左子树由[0, i-1]构成， 其右子树由[i+1, n]构成。
	array[i] = ∑ array[0...k] * [ k+1....i]     0<=k<i-1

	Count[2] = Count[0] * Count[1]   (1为根的情况)
                  + Count[1] * Count[0]  (2为根的情况。

    Count[3] = Count[0]*Count[2]  (1为根的情况)
               + Count[1]*Count[1]  (2为根的情况)
               + Count[2]*Count[0]  (3为根的情况)
*/
func numTrees(n int) int {
	if n <= 1 {
		return 1
	}

	array := make([]int, n+1)
	array[0], array[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			array[i] += array[j-1] * array[i-j]
		}
	}
	return array[n]

}

/*
	不仅要统计数量,还需要把树列举出来.
*/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreesI(1, n)
}

func generateTreesI(left, right int) (ret []*TreeNode) {
	if left > right {
		return []*TreeNode{nil} // 这里注意不要直接返回nil
	}
	if left == right {
		return []*TreeNode{&TreeNode{Val: left}}
	}

	for num := left; num <= right; num++ {
		l := generateTreesI(left, num-1)
		r := generateTreesI(num+1, right)

		for _, vl := range l {
			for _, vr := range r {
				root := &TreeNode{Val: num, Left: vl, Right: vr}
				ret = append(ret, root)
			}
		}
	}
	return ret
}

// 不仅要保证每个节点都是搜索二叉树,而且要保证从上至下都满足二叉树的性质
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	sorted := []*TreeNode{}
	p := root

	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(sorted) > 0 && tmp.Val <= sorted[len(sorted)-1].Val {
				return false
			}
			sorted = append(sorted, tmp)
			p = tmp.Right
		}
	}
	return true
}

// 修复二叉树
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	list := []*TreeNode{}
	vals := []int{}

	list, vals = inorder(root, list, vals)
	sort.Ints(vals)
	for i, v := range list {
		v.Val = vals[i]
	}
}

func inorder(root *TreeNode, list []*TreeNode, vals []int) ([]*TreeNode, []int) {
	if root == nil {
		return list, vals
	}
	list, vals = inorder(root.Left, list, vals)
	list = append(list, root)
	vals = append(vals, root.Val)
	list, vals = inorder(root.Right, list, vals)
	return list, vals
}

// 按层遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	rets := [][]int{}

	for len(stack) > 0 {
		ret := []int{}
		tmpStack := []*TreeNode{}
		for _, v := range stack {

			ret = append(ret, v.Val)
			if v.Left != nil {
				tmpStack = append(tmpStack, v.Left)
			}
			if v.Right != nil {
				tmpStack = append(tmpStack, v.Right)
			}

		}
		rets = append(rets, ret)
		stack = tmpStack
	}
	return rets
}

// 波形按层遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	rets := [][]int{}
	flag := true

	for len(stack) > 0 {
		ret := []int{}
		tmpStack := []*TreeNode{}
		for i := len(stack) - 1; i >= 0; i-- {
			v := stack[i]
			ret = append(ret, v.Val)
			if flag {
				if v.Left != nil {
					tmpStack = append(tmpStack, v.Left)
				}
				if v.Right != nil {
					tmpStack = append(tmpStack, v.Right)
				}
			} else {
				if v.Right != nil {
					tmpStack = append(tmpStack, v.Right)
				}
				if v.Left != nil {
					tmpStack = append(tmpStack, v.Left)
				}
			}
		}
		flag = !flag
		rets = append(rets, ret)
		stack = tmpStack
	}
	return rets
}

// 根据前序便利和中序遍历构造二叉树
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	vals := map[int]int{}
	for i, v := range inorder {
		vals[v] = i
	}

	root := &TreeNode{Val: preorder[0]}
	for _, v := range preorder[1:] {
		for p := root; p != nil; {
			if vals[v] > vals[p.Val] {
				if p.Right != nil {
					p = p.Right
				} else {
					p.Right = &TreeNode{Val: v}
					break
				}
			} else {
				if p.Left != nil {
					p = p.Left
				} else {
					p.Left = &TreeNode{Val: v}
					break
				}
			}
		}
	}
	return root
}

// 根据后序便利和中序遍历构造二叉树
func buildTree3(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	vals := map[int]int{}
	for i, v := range inorder {
		vals[v] = i
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	for i := len(postorder) - 2; i >= 0; i-- {
		v := postorder[i]
		for p := root; p != nil; {
			if vals[v] > vals[p.Val] {
				if p.Right != nil {
					p = p.Right
				} else {
					p.Right = &TreeNode{Val: v}
					break
				}
			} else {
				if p.Left != nil {
					p = p.Left
				} else {
					p.Left = &TreeNode{Val: v}
					break
				}
			}
		}
	}
	return root
}

// 按层打印且从下到上
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	rets := [][]int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		ret := []int{}
		tmpStack := []*TreeNode{}
		for _, v := range stack {
			ret = append(ret, v.Val)

			if v.Left != nil {
				tmpStack = append(tmpStack, v.Left)
			}
			if v.Right != nil {
				tmpStack = append(tmpStack, v.Right)
			}

		}
		rets = append(rets, ret)
		stack = tmpStack
	}

	// 翻转
	count := len(rets) / 2
	if len(rets)%2 == 0 {
		count--
	}
	for i := 0; i <= count; i++ {
		rets[i], rets[len(rets)-1-i] = rets[len(rets)-1-i], rets[i]
	}
	return rets
}

// 给出一个升序排列的数组,转化为平衡二叉树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := (len(nums) - 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

// 是否为平衡二叉树,重点是求节点的高度,匿名函数解法
func isBalanced1(root *TreeNode) bool {
	flag := true
	var help func(rootx *TreeNode) int
	help = func(rootx *TreeNode) int {
		if !flag {
			return -1
		}
		if rootx == nil {
			return 0
		}
		leftHeight := help(rootx.Left)
		rightHeight := help(rootx.Right)
		if (leftHeight-rightHeight) > 1 || (rightHeight-leftHeight) > 1 {
			flag = false
			return -1
		}
		var height int
		if leftHeight > rightHeight {
			height = leftHeight
		} else {
			height = rightHeight
		}
		return 1 + height
	}
	help(root)
	return flag
}

// 是否为平衡二叉树,重点是求节点的高度,非匿名函数
func isBalanced2(root *TreeNode) bool {
	_, flag := TreeHeight(root)
	return flag
}

// 求节点高度
func TreeHeight(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHeight, flag := TreeHeight(root.Left)
	if !flag {
		return -1, false
	}
	rightHeight, flag := TreeHeight(root.Right)
	if !flag {
		return -1, false
	}
	if (leftHeight-rightHeight) > 1 || (rightHeight-leftHeight) > 1 {
		return -1, false
	}
	var height int
	if leftHeight > rightHeight {
		height = leftHeight
	} else {
		height = rightHeight
	}
	return 1 + height, true
}

// 根节点到叶子节点的最小高度值
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftHeight, rightHeight := -1, -1
	if root.Left != nil {
		leftHeight = minDepth(root.Left)
	}
	if root.Right != nil {
		rightHeight = minDepth(root.Right)
	}

	var height int
	if leftHeight == -1 {
		height = rightHeight
	} else if rightHeight == -1 {
		height = leftHeight
	} else if leftHeight < rightHeight {
		height = leftHeight
	} else {
		height = rightHeight
	}
	return 1 + height
}
