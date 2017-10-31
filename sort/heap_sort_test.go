package sort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	nums := []int{16, 7, 3, 20, 17, 8}
	// 20,17,8,7,16,3
	HeapSort(nums)
	if !reflect.DeepEqual(nums, []int{3, 7, 8, 16, 17, 20}) {
		t.Error(nums)
	}
}
