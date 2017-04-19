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

// abcc abcac ohomm
func Substring2(s string) int {
	m := make(map[rune]int, len(s))
	var max, tmp int
	for i, v := range s {
		if j, exist := m[v]; exist {
			if max == 0 {
				max = i // 第一次
			}
			sub := i - j
			if sub > max {
				max = sub
			}
			tmp = sub
			m = make(map[rune]int, len(s))
		} else {
			tmp++
		}
		m[v] = i
	}
	if tmp > max {
		max = tmp
	}
	return max
}
