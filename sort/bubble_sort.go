package sort

import "fmt"

// 冒泡排序
// func BubbleSort(nums []int) {
// 	j := len(nums) - 1
// 	for ; j > 0; j-- {
// 		for i := 0; i < j; i++ {
// 			if nums[i] > nums[i+1] {
// 				nums[i], nums[i+1] = nums[i+1], nums[i]
// 			}
// 		}
// 	}
// }

// 思路 :每次遍历将筛选一个最大的放在尾部,每次遍历之后j--

func BubbleSort(nums []int) {
	j := len(nums) - 1
	for ; j > 0; j-- {
		for i := 0; i < j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	var c = make(chan int, 0)
	fmt.Println(c)
}
