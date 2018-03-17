package sort

// func InsertSort(nums []int) {
// 	for i := 0; i < len(nums)-1; i++ {
// 		n := nums[i+1]
// 		for j := i; j >= 0; j-- {
// 			if nums[j] > n { // 大值都右移,且原值与n交换
// 				nums[j+1], nums[j] = nums[j], n
// 			} else {
// 				nums[j+1] = n
// 				break
// 			}
// 		}
// 	}
// }

// 插入排序
// 思路: 每次将i+1插入到(0,i)中,i++,插入的时候,最好按i -> 0的顺序遍历

// func InsertSort(nums []int) {
// 	for i := 0; i < len(nums)-1; i++ { // [0,i] 表示已排序部分
// 		n := nums[i+1]
// 		j := i
// 		for ; j >= 0 && nums[j] > n; j-- {
// 			nums[j+1] = nums[j]
// 			nums[j] = n
// 		}
// 		nums[j+1] = n
// 	}
// }

// 3 1 2
// 0 4 5
func InsertSort(nums []int) {
	gap := len(nums) / 2
	for gap > 0 {
		for col := 0; col < gap; col++ {

			j := gap
			for j < len(nums) {
				for i := 0; i < gap; i++ {

				}

				j += gap
			}

		}
	}
}
