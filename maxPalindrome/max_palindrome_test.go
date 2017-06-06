package maxPalindrome

import (
	"testing"
)

func TestMax(t *testing.T) {
	in := []string{
		"abcdcba",
		"abc",
		"ccc",
		"babad",
	}
	out := []string{
		"abcdcba",
		"a",
		"ccc",
		"bab",
	}
	for i, _ := range in {
		t.Log(longestPalindrome(in[i]) == out[i])
	}
}

func TestPoint(t *testing.T) {
	in := "aba"
	if l, _, _ := Point(in, 1); l != 3 {
		t.Error("1", l)
	}

	in = "babad"
	if l, _, _ := Point(in, 2); l != 3 {
		t.Error("2", l)
	}
}

func TestEnd(t *testing.T) {
	s := "ccc"
	t.Log(getEnd(s, 0))
	s = "abc"
	t.Log(getEnd(s, 0))
	s = "babad"
	t.Log(getEnd(s, 2))
}
