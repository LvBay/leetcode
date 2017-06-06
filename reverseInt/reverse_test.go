package reverseInt

import (
	"testing"
)

func TestW(t *testing.T) {
	if m := maxm(1000); m != 3 {
		t.Error("1", m)
	}
	if m := maxm(2000); m != 3 {
		t.Error("2", m)
	}
	if m := maxm(123); m != 0 {
		t.Error("3", m)
	}
}

func TestReserve(t *testing.T) {
	if m := reverse(-123); m != -321 {
		t.Error("1", m)
	}
	if m := reverse(0); m != 0 {
		t.Error("2", m)
	}
	if m := reverse(9999999999); m != 0 {
		t.Error("2", m)
	}
}
