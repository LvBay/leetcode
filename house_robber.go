package leetcode

func rob(nums []int) int {
	var rob, noRob int
	for i := 0; i < len(nums); i++ {
		tmp := noRob + nums[i]
		noRob = max(rob, noRob)
		rob = tmp
	}
	return max(rob, noRob)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
	//     2 2 4 4 1
	// no  0 2 2 6 6
	// y   2 0 6 6 7
}
