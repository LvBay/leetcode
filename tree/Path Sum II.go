package main

// 查制指定目标值的路径
func pathSum(root *TreeNode, sum int) [][]int {

	ret := [][]int{}
	old := []int{}
	ret = help(root, 0, sum, old, ret)
	return ret
}

func help(root *TreeNode, sum, target int, old []int, ret [][]int) [][]int {
	if root == nil {
		return ret
	}

	hist := append(old, root.Val)
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
