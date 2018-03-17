package tree

import (
	"reflect"
	"testing"
)

func TestCreateBstree(t *testing.T) {
	in := [][]int{
		[]int{1},
		[]int{3, 1, 5},
		[]int{1, 3, 5},
		[]int{10, 7, 6, 15, 9, 13, 14},
	}
	out := []*TreeNode{
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}

	for i := 0; i < len(in); i++ {
		get := CreateBstree(in[i])
		if !compare(get, out[i]) {
			t.Error(i)
			return
		}
	}
}

func TestInOrder(t *testing.T) {
	in := []*TreeNode{
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}
	out := [][]int{
		[]int{1},
		[]int{1, 3, 5},
		[]int{1, 3, 5},
		[]int{6, 7, 9, 10, 13, 14, 15},
	}

	for i := 0; i < len(in); i++ {
		get := in[i].InOrder()
		if !compare(get, out[i]) {
			t.Error(i)
			return
		}
	}
}

func TestPreOrder(t *testing.T) {
	in := []*TreeNode{
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}
	out := [][]int{
		[]int{1},
		[]int{3, 1, 5},
		[]int{1, 3, 5},
		[]int{10, 7, 6, 9, 15, 13, 14},
	}

	for i := 0; i < len(in); i++ {
		get := in[i].PreOrder()
		if !compare(get, out[i]) {
			t.Error(i)
			return
		}
	}
}

func TestPostOrder(t *testing.T) {
	in := []*TreeNode{
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}
	out := [][]int{
		[]int{1},
		[]int{1, 5, 3},
		[]int{5, 3, 1},
		[]int{6, 9, 7, 14, 13, 15, 10},
	}

	for i := 0; i < 4; i++ {
		get := in[i].PostOrder()
		if !compare(get, out[i]) {
			t.Error(i)
			return
		}
	}
}

func TestLevel(t *testing.T) {
	in := []*TreeNode{
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}
	out := [][][]int{
		[][]int{
			[]int{1},
		},
		[][]int{
			[]int{3},
			[]int{1, 5},
		},
		[][]int{
			[]int{1},
			[]int{3},
			[]int{5},
		},
		[][]int{
			[]int{10},
			[]int{7, 15},
			[]int{6, 9, 13},
			[]int{14},
		},
	}

	for i := 0; i < len(in); i++ {
		get := in[i].Level()
		if !compare(get, out[i]) {
			t.Error(i)
			return
		}
	}
}

func TestFix(t *testing.T) {
	in := []*TreeNode{
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 13,
				Left: &TreeNode{
					Val: 15,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}
	out := []*TreeNode{
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}

	for i := 0; i < len(in); i++ {
		get := in[i].Fix()
		if !compare(get, out[i]) {
			t.Error(i)
			return
		}
	}
}

func TestAdd(t *testing.T) {
	in := []*TreeNode{
		nil,
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
				},
			},
		},
	}
	in2 := []int{1, 5, 5, 14}
	out := []*TreeNode{
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}
	for i := 0; i < 1; i++ {
		get := in[i].Add(in2[i])
		if !compare(get, out[i]) {
			t.Error(i)
			return
		}
	}
}

func TestParent(t *testing.T) {
	in := []*TreeNode{
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}
	in2 := []int{1, 1, 5, 6}
	out := []int{0, 3, 3, 7}
	for i := 0; i < len(in); i++ {
		get := in[i].Parent(in2[i])
		if i == 0 {
			if get != nil {
				t.Error(i)
				return
			}
		} else if get.Val != out[i] {
			t.Error(i)
		}
	}
}

func TestFind(t *testing.T) {
	in := []*TreeNode{
		&TreeNode{Val: 1},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}
	in2 := []int{1, 1, 5, 6}
	out := []int{1, 1, 5, 6}
	for i := 0; i < len(in); i++ {
		get := in[i].Find(in2[i])
		if get.Val != out[i] {
			t.Error(i)
		}
	}
}

func TestDelete(t *testing.T) {
	in := []*TreeNode{
		&TreeNode{Val: 1}, // 删除
		&TreeNode{
			Val: 3, // 删除
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3, // 删除
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15, // 删除
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7, // 删除
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}
	in2 := []int{1, 3, 3, 15, 7}
	out := []*TreeNode{
		nil,
		&TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 1,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 13,
				Right: &TreeNode{
					Val: 14,
				},
			},
		},
		&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
		},
	}
	for i := 0; i < 4; i++ {
		get := in[i].Delete(in2[i])
		if !compare(get, out[i]) {
			t.Error(i)
		}
	}
}

// 比较两个tree值是否相等
func compare(get, want interface{}) bool {
	// utils.Display("get", get)
	// utils.Display("want", want)
	ret := reflect.DeepEqual(get, want)
	return ret
}

// 比较两个tree值是否相等
func compare1(get, want *TreeNode) bool {
	if get == nil {
		return want == nil
	}
	if want == nil {
		return get == nil
	}

	if get.Val != want.Val {
		return false
	}
	if !compare(get.Left, want.Left) {
		return false
	}
	if !compare(get.Right, want.Right) {
		return false
	}
	return true
}

func compare2(get, want []int) bool {
	if len(get) != len(want) {
		return false
	}
	for i, v := range get {
		if v != want[i] {
			return false
		}
	}
	return true
}
