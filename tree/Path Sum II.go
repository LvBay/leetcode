package tree

// 查制指定目标值的路径
func pathSum(root *TreeNode, sum int) [][]int {

	ret := [][]int{}
	collected := []int{}
	ret = help(root, 0, sum, collected, ret)
	return ret
}

func help(root *TreeNode, sum, target int, collected []int, ret [][]int) [][]int {

	if root == nil {
		return ret
	}

	// don't user hist := append(collected,root.Val)
	hist := make([]int, len(collected))
	copy(hist, collected)
	hist = append(hist, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val+sum == target {
			ret = append(ret, hist)

		}
		return ret
	}
	sum += root.Val
	ret = help(root.Left, sum, target, hist, ret)

	ret = help(root.Right, sum, target, hist, ret)
	return ret
}
