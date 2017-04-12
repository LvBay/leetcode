package linkNode

import (
	"encoding/json"
	// "fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方便查看数据
func ToString(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

func (list *ListNode) GetRear() (rear *ListNode) {
	if list == nil {
		return nil
	}
	// only one Listnode
	if list.Next == nil {
		return list
	}
	rear = list.Next
	for rear.Next != nil {
		rear = rear.Next
	}
	return rear
}

func (list *ListNode) Append(data *ListNode) *ListNode {
	if list == nil {
		return data
	}
	rear := list.GetRear()
	rear.Next = data
	return list
}

func (list1 *ListNode) Add(list2 *ListNode) *ListNode {
	list3 := &ListNode{}
	isHead := true
	var carry int
	for list1 != nil || list2 != nil {
		var sum int
		if list1 != nil {
			sum += list1.Val
			list1 = list1.Next
		}
		if list2 != nil {
			sum += list2.Val
			list2 = list2.Next
		}
		sum += carry
		carry = 0
		if sum >= 10 {
			sum -= 10
			carry = 1
		}
		tmp := &ListNode{sum, nil}
		if isHead {
			list3 = tmp
			isHead = false
		} else {
			list3.Append(tmp)
		}
	}
	return list3
}

// leetcode答案
func AddTwoList(l1, l2 *ListNode) *ListNode {
	list3 := &ListNode{}
	rear := &ListNode{}
	isHead := true
	var carry int
	for l1 != nil || l2 != nil || carry > 0 {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = 0
		if sum >= 10 {
			sum -= 10
			carry = 1
		}
		tmp := &ListNode{sum, nil}
		if isHead {
			list3 = tmp
			rear = tmp
			isHead = false
		} else {
			// 运用指针
			rear.Next = tmp // 更新链表尾部值和地址
			rear = tmp      // 便于下次从尾部追加数据
		}

	}
	return list3
}
