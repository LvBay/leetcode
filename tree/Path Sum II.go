package tree

import "fmt"

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
		fmt.Println("叶子节点:", root.Val)
		if root.Val+sum == target {
			ret = append(ret, hist)
			fmt.Println("ret1:", ret)
		}

		return ret
	}
	sum += root.Val
	ret = help(root.Left, sum, target, hist, ret)
	fmt.Println("ret2:", ret)

	ret = help(root.Right, sum, target, hist, ret)
	fmt.Println("ret3:", ret)
	return ret
}
