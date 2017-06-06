package palindrome_number

import (
// "fmt"
)

func GetArray(x int) (arr []int) {
	if x < 0 {
		x = 0 - x
	}
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	return arr
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} // 不明白为什么leet上判定所有负数都不通过
	arr := GetArray(x)
	length := len(arr)
	i, j := 0, length-1
	for j >= i {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
