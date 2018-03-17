package sort

func CombineSort(nums []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		if right-left > 1 {
			CombineSort(nums, left, mid)
			CombineSort(nums, mid+1, right)
		}
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {

	i := left
	j := mid + 1
	tmp := make([]int, 0, right-left+1)
	for {
		if i > mid && j > right {
			break
		}
		if i > mid {
			tmp = append(tmp, nums[j])
			j++
			continue
		}

		if j > right {
			tmp = append(tmp, nums[i])
			i++
			continue
		}

		if nums[i] > nums[j] {
			tmp = append(tmp, nums[j])
			j++
		} else {
			tmp = append(tmp, nums[i])
			i++
		}
	}
	for i := 0; i < len(tmp); i++ {
		nums[left+i] = tmp[i]
	}
	// beego.Info(tmp)
	// beego.Info(nums)
}
