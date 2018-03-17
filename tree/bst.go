package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	P     *TreeNode
}

import (
	"sort"
)

func CreateBstree(nums []int) (root *TreeNode) {
	if len(nums) == 0 {
		return nil
	}
	var curr *TreeNode
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			root = &TreeNode{Val: nums[0]}
			continue
		}
		curr = root
		for curr != nil {
			if nums[i] > curr.Val {
				if curr.Right == nil {
					curr.Right = &TreeNode{Val: nums[i]}
					break
				} else {
					curr = curr.Right
				}
			} else {
				if curr.Left == nil {
					curr.Left = &TreeNode{Val: nums[i]}
					break
				} else {
					curr = curr.Left
				}
			}
		}
	}
	// utils.Display("root:", root)
	return root
}

// 中序遍历 非递归
func (root *TreeNode) InOrder() []int {
	stack := []*TreeNode{}
	curr := root
	ret := []int{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, tmp.Val)
			curr = tmp.Right
		}
	}
	// utils.Display("ret:", ret)
	return ret
}

// 中序遍历 非递归
func (root *TreeNode) PreOrder() []int {
	stack := []*TreeNode{}
	curr := root
	ret := []int{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			ret = append(ret, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}
		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curr = tmp.Right
		}
	}
	// utils.Display("ret:", ret)
	return ret
}

// 后序遍历 非递归
// 应用场景:表达式计算
func (root *TreeNode) PostOrder() []int {
	stack := []*TreeNode{root}
	curr := root
	var pre *TreeNode
	ret := []int{}
	for len(stack) > 0 {
		curr = stack[len(stack)-1]
		// 如果栈顶是叶子节点或者上个节点等于当前节点的子节点
		// pre != nil 是为了防止1,3,4这样单向树的情况
		// curr.Left == pre || curr.Right == pre 是为了防止从下向上后,因为父节点不是子节点,所以会陷入死循环
		if curr.Left == nil && curr.Right == nil || (pre != nil && (curr.Left == pre || curr.Right == pre)) {
			ret = append(ret, curr.Val)
			stack = stack[:len(stack)-1]
			pre = curr
		} else {
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
		}
	}
	// utils.Display("ret:", ret)
	return ret
}

// 按层打印(bfs)
func (root *TreeNode) Level() [][]int {
	stack := []*TreeNode{root}
	ret := [][]int{}
	count := 1
	for len(stack) > 0 {
		tmpRet := make([]int, 0, count)
		tmpStack := make([]*TreeNode, 0, 2*len(stack))
		for _, v := range stack {
			tmpRet = append(tmpRet, v.Val)
			if v.Left != nil {
				tmpStack = append(tmpStack, v.Left)
			}
			if v.Right != nil {
				tmpStack = append(tmpStack, v.Right)
			}
		}
		ret = append(ret, tmpRet)
		stack = tmpStack
		count *= 2
	}
	// utils.Display("ret:", ret)
	return ret
}

// 添加节点
func (root *TreeNode) Add(n int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: n}
		return root
	}
	curr := root
	for curr != nil {
		if n > root.Val {
			if root.Right == nil {
				root.Right = &TreeNode{Val: n}
				break
			} else {
				curr = curr.Right
			}
		} else {
			if root.Left == nil {
				root.Left = &TreeNode{Val: n}
				break
			} else {
				curr = curr.Left
			}
		}
	}
	return root
}

// 删除节点
// 1. 被删除节点z没有任何子女。是最简单的情况，将z父节点12的儿子指针置为null，再free掉节点z就好;
// 2. 被删节点z有且只有有一个子女。这个情况稍微麻烦点，但说起来比较容易，将节点z的父节点15和其唯一子节点20连接起来就好，然后free掉节点z就好;
// 3. 被删节点z有两个子女。这个是最复杂的，但是其基本思想还是没变，将第三种情况转化为第二种情况处理即可。那么怎么做呢？解决方法就是去找z的后继（后继是什么？就是比节点z中关键字5大的所有节点中最小的那个）。
// 要删除有两个子女的节点z，就去找它的后继节点h，它的后继节点h能保证只有一个右儿子（为什么？如果它的后继节点有左儿子，那么节点y就不可能是z的后继节点。算导书上p155 练习 12.2-5有要求证明，不懂的可以稍微思考下）。
// 基于上述事实，我们找到了z的后继节点h，那么就已经将情况3转化为了情况2了，接下来要做的操作就是将h的父节点hp与h的(右)子节点连起来。
// 这时还没做完，因为z才是待删除的节点，我们要将节点z出的关键字替换为节点h的关键字，如图中所示。然后free掉节点h。
func (root *TreeNode) Delete(n int) *TreeNode {
	fake := &TreeNode{Left: root}
	z := root.Find(n)
	var p *TreeNode
	if z.Left == nil && z.Right == nil {
		p = fake.Parent(z.Val)
		// 通过父节点,删除该节点
		if p.Left == z { // 该节点为左子节点
			p.Left = nil
		} else if p.Right == z { // 该节点为右子节点
			p.Right = nil
		}
		return fake.Left
	}
	// z节点有一个右节点
	if z.Left == nil && z.Right != nil {
		p = fake.Parent(z.Val)
		if p.Left == z { // 该节点为左子节点
			p.Left = z.Right
		} else if p.Right == z { // 该节点为右子节点
			p.Right = z.Right
		}
		return fake.Left
	}

	// z节点有一个左节点
	if z.Left != nil && z.Right == nil {
		p = fake.Parent(z.Val)
		if p.Left == z { // 该节点为左子节点
			p.Left = z.Left
		} else if p.Right == z { // 该节点为右子节点
			p.Right = z.Left
		}
		return fake.Left
	}
	// z节点有两个节点
	h := z.Successor()     // z节点的后继节点h,注意h节点只会存在右子节点
	p = fake.Parent(h.Val) // 后继节点的父节点 hp
	if p.Left == h {       // 后继节点h是父节点的左节点
		p.Left = h.Right
	} else if p.Right == h { // 后继节点h是父节点的右节点
		p.Right = h.Right
	}
	z.Val = h.Val

	return fake.Left
}

// 查询节点的后继 假设n肯定有右孩子
func (root *TreeNode) Successor() *TreeNode {
	curr := root.Right
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

// 查询节点的父节点
func (root *TreeNode) Parent(n int) *TreeNode {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if tmp.Left != nil && tmp.Left.Val == n {
				return tmp
			}
			if tmp.Right != nil && tmp.Right.Val == n {
				return tmp
			}
			curr = tmp.Right
		}
	}
	return nil
}

// 修复bst
func (root *TreeNode) Fix() *TreeNode {
	stack := []*TreeNode{}
	curr := root
	ret := []int{}
	retNodes := []*TreeNode{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, tmp.Val)
			retNodes = append(retNodes, tmp)
			curr = tmp.Right
		}
	}
	sort.Ints(ret)
	for i, v := range ret {
		retNodes[i].Val = v
	}
	return root
}

// 查询节点
func (root *TreeNode) Find(n int) *TreeNode {
	stack := []*TreeNode{}
	curr := root
	ret := []int{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, tmp.Val)
			if tmp.Val == n {
				return tmp
			}
			curr = tmp.Right
		}
	}
	return nil
}
