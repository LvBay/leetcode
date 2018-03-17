package array

// 从nums中插入x
func insert(nums []int, x int) []int {
	nums = append(nums, x)
	var i int
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] > x {
			nums[i+1] = nums[i]
		} else {
			nums[i+1] = x
			break
		}
	}
	if i == -1 {
		nums[0] = x
	}
	return nums
}
