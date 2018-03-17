package leetcode

func generate(num int) [][]int {
	nums := make([][]int, num)
	for i, _ := range nums {
		nums[i] = make([]int, i+1)
	}
	// l := num/2 + 1
	for i := 0; i < num; i++ {
		for j := 0; j < i+1; j++ {
			if i == 0 || j == 0 || j == i {
				nums[i][j] = 1
			} else {
				nums[i][j] = nums[i-1][j-1] + nums[i-1][j]
			}
		}
	}
	return nums
}
