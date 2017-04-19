package substring_test

import (
	"leetcode/substring"
	"testing"
)

func TestSubstring(t *testing.T) {
	// if s := substring.Substring("bbb"); s != 1 {
	// 	t.Errorf("err 1 want 1, but get %d", s)
	// } else {
	// 	t.Log("PASS1")
	// }
	// if s := substring.Substring("abc"); s != 3 {
	// 	t.Errorf("err 2 want 3, but get %d", s)
	// } else {
	// 	t.Log("PASS2")
	// }
	t.Log(substring.Substring("abc"))
}

func TestRepeat(t *testing.T) {
	if substring.Repeat("abc") != false {
		t.Errorf("err 2-1 ,abc get true")
	} else {
		t.Log("PASS 2-1")
	}
	if substring.Repeat("bb") != true {
		t.Errorf("err 2-2 ,bb get false	")
	} else {
		t.Log("PASS 2-2")
	}
}

func TestSubstring2(t *testing.T) {
	s := []string{"abcdd", "abccb", "abccdfg", "abccdb", "ohomm"}
	res := []int{4, 3, 4, 3, 3}
	for i, v := range s {
		t.Log("i:", i)
		if substring.Substring2(v) != res[i] {
			t.Errorf("want %d, but get %d", res[i], substring.Substring2(v))
		} else {
			t.Logf("%s want %d, but get %d", s[i], res[i], substring.Substring2(v))
		}
	}
}
