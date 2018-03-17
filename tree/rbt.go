package tree

// 红黑树

type RBTree struct {
	Left  *RBTree
	Right *RBTree
	Val   int
	Color string
}

func (root *RBTree) Lrotate() {

}
