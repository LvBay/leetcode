package leetcode

func isUgly(num int) bool {
	nums := []int{2, 3, 5}
	for num > 1 || num < -1 {
		flag := false
		for i := 0; i < len(nums); i++ {
			if num%nums[i] == 0 {
				num = num / nums[i]
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
