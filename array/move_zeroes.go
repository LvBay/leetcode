package array

import (
	"fmt"

	"github.com/astaxie/beego"
)

func moveZeroes(nums []int) {
	var i, j int
	for ; i < len(nums) && j < len(nums); i++ {
		if nums[i] == 0 {
			if j < i {
				j = i
			}
			for ; j < len(nums); j++ {
				if nums[j] != 0 {
					beego.Info(nums)
					beego.Info(fmt.Sprintf("nums[%d]=nums[%d],%d,%d", i, j, nums[i], nums[j]))
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
	beego.Info("ret:", nums)
}
