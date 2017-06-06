package maxPalindrome

func Point(s string, i int) (sub, j, k int) {
	j = i
	k = getEnd(s, i)
	j, k = extend(s, j, k)
	return k - j + 1, j, k
}

func longestPalindrome(s string) string {
	var max, start, end int
	for i := 0; i < len(s); i++ {
		j := i                 // extend左起点
		k := getEnd(s, i)      // extend的右起点
		j, k = extend(s, j, k) // 向两边扩张
		if k-j+1 > max {
			max = k - j + 1
			start = j
			end = k
		}
	}
	return s[start : end+1]
}

func getEnd(s string, i int) (k int) {
	k = i
	for k < len(s)-1 && s[k+1] == s[i] {
		k++
	}
	return k
}

func extend(s string, j, k int) (jj, kk int) {
	jj, kk = j, k
	for jj > 0 && kk < len(s)-1 && s[jj-1] == s[kk+1] {
		jj--
		kk++
	}
	return jj, kk
}
