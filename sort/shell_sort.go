package sort

// 1 3 4
// 2 5 6
func ShellSort(nums []int) {
	n := len(nums)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ { // i=3,4,5
			tmp := nums[i] // 2
			j := i         // 3
			for j >= gap && nums[j-gap] > tmp {
				nums[j] = nums[j-gap]
				j = j - gap
			}
			nums[j] = tmp
		}
		gap = gap / 2
	}
}
