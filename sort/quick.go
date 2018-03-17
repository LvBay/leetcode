// package sort
package main

import (
	"github.com/astaxie/beego"
)

func quickSort(nums []int) {
	help2(nums, 0, len(nums)-1)
}

// 双指针i,j   i从左至右寻找大于基准的值,j从右至左寻找小于基准的值
// 因为最后一步,需要交换基准和nums[i],所以要先执行j,再执行i,保证最后的nums[i]是小于基准值
func help(nums []int, start, end int) {
	if start >= end {
		return
	}
	beego.Info("start,end:", start, end)
	pivot := nums[start]
	i := start
	j := end
	for i < j {

		for ; nums[j] >= pivot && i < j; j-- {

		}

		for ; nums[i] <= pivot && i < j; i++ {

		}

		if i == j {
			beego.Info("i==j", i, j)
			nums[start], nums[i] = nums[i], nums[start]
			break
		} else {
			beego.Info("i!=j", i, j)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	help(nums, start, i-1)
	help(nums, i+1, end)
}

// 3 1 4
func help2(nums []int, start, end int) {
	if start >= end {
		return
	}
	beego.Info("start,end:", start, end)
	pivot := nums[end]
	i := start
	j := end
	for i < j {

		for ; nums[i] <= pivot && i < j; i++ {

		}

		for ; nums[j] >= pivot && i < j; j-- {

		}

		if i == j {
			beego.Info("i==j", i, j)
			nums[end], nums[i] = nums[i], nums[end]
			break
		} else {
			beego.Info("i!=j", i, j)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	help2(nums, start, i-1)
	help2(nums, i+1, end)
}

// 分区
// 挑选right为基准pivot,变量storeIndex为从左到右小于povit的个数,初始值为left
// 从左到右(right-1)遍历,如果array[i]小于基准,移到左边(根据storeIndex)
// 最后将array[right]移到storeIndex位置,保证storeIndex左边都是小于基准的值,右边都是大于基准的值
func partition(array []int, left, right int) {
	if left >= right {
		return
	}
	var storeIndex int = left
	var pivot int = array[right] // 直接选最右边的元素为基准元素
	for i := left; i < right; i++ {
		if array[i] < pivot {
			array[storeIndex], array[i] = array[i], array[storeIndex]
			storeIndex++ // 交换位置后，storeIndex 自增 1，代表下一个可能要交换的位置
		}
	}
	array[right], array[storeIndex] = array[storeIndex], array[right] // 将基准元素放置到最后的正确位置上
	partition(array, left, storeIndex-1)
	partition(array, storeIndex+1, right)
}

func main() {
	nums := []int{3, 2, 13}
	partition(nums, 0, len(nums)-1)
	beego.Info(nums)
}

func quickSort2(nums []int) {

}
