package array

import (
	"github.com/astaxie/beego"
)

func containsDuplicate(nums []int) bool {
	var max, min int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			max = nums[i]
		} else {
			if nums[i] > max {
				max = nums[i]
			}
			if nums[i] < min {
				min = nums[i]
			}
		}
	}
	arr := make([]int, max-min+1)
	beego.Info("len arr:", len(arr))
	for i := 0; i < len(nums); i++ {
		if arr[nums[i]-min] == 0 {
			arr[nums[i]-min] = 1
		} else {
			return true
		}
	}
	return false
}
