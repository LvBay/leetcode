package tree

import (
	"fmt"
)

func sumNumbers(root *TreeNode) int {
	collected := []int{}
	ret := [][]int{}
	ret = helpSumNumbers(root, collected, ret)
	sum := 0
	for _, v := range ret {
		sum += Sum(v)
	}
	return sum
}

func helpSumNumbers(root *TreeNode, collected []int, ret [][]int) [][]int {
	if root == nil {
		return ret
	}
	hist := make([]int, len(collected))
	copy(hist, collected)
	hist = append(hist, root.Val)
	if root.Left == nil && root.Right == nil {
		ret = append(ret, hist)
	}
	ret = helpSumNumbers(root.Left, hist, ret)
	ret = helpSumNumbers(root.Right, hist, ret)
	return ret
}

func Sum(nums []int) int {
	sum := 0
	for i, _ := range nums {
		m := 10
		if i == 0 {
			m = 1
		} else {
			for j := i; j > 1; j-- {
				m *= 10
			}
		}
		sum += m * nums[len(nums)-1-i]
	}
	return sum
}
