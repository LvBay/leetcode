package reverseInt

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var n int
	var flag bool // 负数

	if x < 0 {
		x = 0 - x
		flag = true
	}

	for i := maxm(x); i >= 0; i-- {
		rm := x % 10 // 余数
		n += pow10(i) * rm
		fmt.Println("rm:", rm)
		x /= 10
	}

	if flag {
		n = 0 - n
	}
	// overflow 2147483647
	if n > math.MaxInt32 || n < math.MinInt32 {
		n = 0
	}
	return n
}

// 10的最高次幂
func maxm(x int) (i int) {
	for {
		m := x / 10
		if m <= 0 {
			break
		}
		i++
		x /= 10
	}
	return i
}

// 10的x次幂
func pow10(x int) int {
	y := 1
	for i := 0; i < x; i++ {
		y *= 10
	}
	return y
}
