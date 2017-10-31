package sort

func HeapSort(nums []int) {
	length := len(nums)
	buildHeap(nums)
	for i := 1; i < len(nums); i++ {
		nums[0], nums[length-i] = nums[length-i], nums[0]
		ajustHeap(nums[:length-i], 0)
	}
}

func buildHeap(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		ajustHeap(nums, i)
	}
}

func ajustHeap(nums []int, pos int) {
	length := len(nums)
	for j := pos; j < length; {
		left := 2*j + 1  // 左节点
		right := 2*j + 2 // s右节点

		if left >= length {
			break // 不存在子节点
		}
		max := left
		if right < length && nums[right] > nums[left] {
			max = right
		}

		if nums[max] > nums[j] { // 当前节点小于子节点,交换
			nums[j], nums[max] = nums[max], nums[j]
			j = max
		} else {
			break
		}
	}
}
