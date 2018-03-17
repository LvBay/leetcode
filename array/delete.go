package array

// 从nums中删除x元素,假设x一定存在于nums中
func delete(nums []int, x int) []int {
	var i, j, count int
	for i < len(nums)-count && j < len(nums) {
		for ; j < len(nums) && nums[j] == x; j++ {
			count++
		}
		if j >= len(nums) {
			break
		}
		if i != j {
			nums[i] = nums[j]
		}
		i++
		j++
	}
	// beego.Info(nums[:len(nums)-count])
	return nums[:len(nums)-count]
}
