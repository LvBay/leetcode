package sort

import (
	"testing"

	"github.com/astaxie/beego"
)

func TestBubbleSort(t *testing.T) {
	in := [][]int{
		[]int{2, 1},
		[]int{2, 1, 3, 5, 0},
	}
	for i := 0; i < len(in); i++ {
		BubbleSort(in[i])
		beego.Info(in[i])
	}
}
