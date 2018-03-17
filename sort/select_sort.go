package sort

// // 选择排序
// func SelectSort(nums []int) {
// 	for i := 0; i < len(nums); i++ {
// 		min := nums[i]
// 		tmp := i
// 		for j := i; j < len(nums); j++ {
// 			if nums[j] < min {
// 				tmp = j
// 				min = nums[j]
// 			}
// 		}
// 		if tmp != i {
// 			nums[tmp], nums[i] = nums[i], nums[tmp]
// 		}
// 	}
// }

// 选择排序
// 思路: 每次从(i,j)中选出一个最小值,放在(0,i)后面

func SelectSort(nums []int) {

	for i := 0; i < len(nums); i++ {
		min := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}
