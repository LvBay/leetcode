package tree

func TestLrotate(t *testing.T) {
	in := []*RBTree{
		&RBTree{
			Val: 10,
			Left: &RBTree{
				Val: 8,
			},
			Right: &RBTree{
				Val: 16,
				Left: &RBTree{
					Val: 14,
				},
				Right: &RBTree{
					Val: 18,
				},
			},
		},
	}
	in2 := []int{10},
out := &RBTree{
		Val:16,
		Left:&RBTree{

		},
	},
}
