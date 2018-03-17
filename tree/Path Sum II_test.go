package tree

import "testing"

func TestPathSum(t *testing.T) {
	root := buildNomalTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1})
	ret := pathSum(root, 22)
	t.Log(ret)
}
