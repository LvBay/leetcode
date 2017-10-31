package linkNode

import (
	"testing"
	// "fmt"
)

func TestGetRear(t *testing.T) {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n2.Next = n3
	n1.Next = n2
	if n3.GetRear() != n3 {
		t.Error("err1:", n3.GetRear())
	} else {
		t.Log("PASS1")
	}
	if n1.GetRear() != n3 {
		t.Error("want n3,but get %v", n1.GetRear())
	} else {
		t.Log("PASS2")
	}
}

func TestAppend(t *testing.T) {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n1.Append(n2)
	if n1.GetRear() != n2 {
		t.Errorf("err2-1 want n2,but get %v", n2)
	} else {
		t.Log("PASS2-1")
	}
	n1.Append(n3)
	if n1.GetRear() != n3 {
		t.Errorf("err2-2 want n2,but get %v", n3)
	} else {
		t.Log("PASS2-2")
	}
}

func TestAdd(t *testing.T) {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}
	list1 := n1.Append(n2)
	list2 := n3.Append(n4)
	list3 := list1.Add(list2)
	if list3 == nil {
		t.Error("err 3-0 :list3 == nil")
		return
	}
	if list3.Val != 4 {
		t.Errorf("err 3-1 want 4,but get %d", list3.Val)
	} else {
		t.Log("PASS 3-1")
	}

	if list3.Next == nil {
		t.Error("err 3-2 :list3.Next == nil")
		return
	} else {
		t.Log("PASS 3-2")
	}
	if list3.Next.Val != 6 {
		t.Errorf("err 3-3 want 6,but get %d", list3.Next.Val)
	} else {
		t.Log("PASS 3-3")
	}
}

func TestAddTwoList(t *testing.T) {
	n1 := ListNode{2, nil}
	n2 := ListNode{4, nil}
	n3 := ListNode{3, nil}
	n4 := ListNode{5, nil}
	n5 := ListNode{6, nil}
	n6 := ListNode{4, nil}
	list1 := (&n1).Append(&n2).Append(&n3)
	list2 := (&n4).Append(&n5).Append(&n6)
	list3 := AddTwoList(list1, list2)
	if list3.Val != 7 {
		t.Errorf("err 4-1 want 7, get %d", list3.Val)
	} else {
		t.Log("PASS 4-1")
	}
	if list3.Next.Val != 0 {
		t.Errorf("err 4-2 want 0, get %d", list3.Next.Val)
	} else {
		t.Log("PASS 4-2")
	}
	if list3.Next.Next.Val != 8 {
		t.Errorf("err 4-3 want 8, get %d", list3.Next.Next.Val)
	} else {
		t.Log("PASS 4-3")
	}
	n7 := ListNode{5, nil}
	n8 := ListNode{5, nil}
	list3 = AddTwoList(&n7, &n8)
	if list3.Next.Val != 1 {
		t.Errorf("err 4-4 want 1, get %d", list3.Next.Val)
	} else {
		t.Log("PASS 4-4")
	}
}

func TestRevert(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	n5 := Revert(node1)
	for i := 4; i > 0; i-- {
		if i != n5.Val {
			t.Errorf("want %d,bug get %d\n", i, n5.Val)
			return
		}
		n5 = n5.Next
	}
}

func TestCombine(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node3
	node2.Next = node4

	ret := Combine(node1, node2)
	for i := 1; i < 5; i++ {
		if i != ret.Val {
			t.Errorf("want %d,bug get %d\n", i, ret.Val)
			return
		}
		ret = ret.Next
	}
}
