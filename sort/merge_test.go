package sort

import (
	"testing"

	"github.com/astaxie/beego"
)

func TestMerge(t *testing.T) {
	in := []int{5, 2}
	merge(in, 0, 0, 1)
	// beego.Info(in)
}

func TestCombineSort(t *testing.T) {
	in := [][]int{
		[]int{5, 2},
		[]int{1, 3, 2, 6},
		[]int{1, 2, 3, 6},
		[]int{6, 3, 1, 2},
		[]int{3, 1, 2},
		[]int{3, 2, 1},
	}
	for i := 0; i < len(in); i++ {
		CombineSort(in[i], 0, len(in[i])-1)
		beego.Info(in[i])
	}

}
