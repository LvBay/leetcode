package sort

import (
	"testing"

	"github.com/astaxie/beego"
)

func TestInsertSort(t *testing.T) {
	in := [][]int{
		[]int{2, 1},
		[]int{2, 1, 3, 5, 0, 4},
	}
	for i := 0; i < len(in); i++ {
		InsertSort(in[i])
		beego.Info(in[i])
	}
}
