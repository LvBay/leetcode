package leetcode

func reverseVowels(s string) string {
	arr := make([]int, len(s))
	for i, v := range s {

	}
	rs := []byte(s)
	for i, v := range arr {
		rs[v], rs[arr[len(arr)-i]] = rs[arr[len(arr)-i]], rs[v]
	}
	return string(rs)

	switch s 
case "a":
	
}
