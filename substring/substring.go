package substring

import (
	"fmt"
)

func Substring(str string) int {
	fmt.Println("strrrr:", str)
	length := len(str)
	for i := length; i > 0; i-- {
		for j := 0; j+i < length; j++ { // 开始下标
			fmt.Println(str[j : i+j])
			fmt.Println("repeat?:", !Repeat(str[j:i+j]))
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
