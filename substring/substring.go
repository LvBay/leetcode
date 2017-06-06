package substring

import (
// "fmt"
)

func Substring(str string) int {
	length := len(str)
	for i := length; i > 0; i-- {
		for j := 0; j+i <= length; j++ { // 开始下标
			if !Repeat(str[j : i+j]) {
				return i
			}
		}
	}

	return 0
}

func Repeat(str string) bool {
	m := map[rune]bool{}
	for _, b := range str {
		if m[b] == true {
			return true
		}
		m[b] = true
	}
	return false
}

// abcc abccb ohomm abccdfg
func Substring2(s string) int {
	m := map[rune]int{}
	var max, begin int
	for i, v := range s {
		if j, exist := m[v]; exist {
			begin = Max(j+1, begin)
		}
		m[v] = i
		max = Max(max, i-begin+1)
	}
	return max
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
