package linkNode2

import (
	"encoding/json"
	"fmt"
)

// 方便查看数据
func ToString(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

type Student struct {
	Val  int
	Next *Student
}

func Create(x int) (s Student) {
	var head, rear, tmp *Student
	// tmp := &Student{}
	// rear := &Student{}

	for i := 0; i < x; i++ {
		tmp = &Student{Val: i}
		if i == 0 {
			// s = tmp
			head = tmp
			rear = tmp
		} else {
			rear.Next = tmp
			rear = tmp
		}
	}
	fmt.Println("创建链表:", ToString(head))
	return *head
}

func (s *Student) Shift() Student {
	s = s.Next
	fmt.Println("删除头结点:", ToString(s))
	return *s
}

func (s *Student) Del(x int) Student {
	count := 0
	for s.Next != nil {
		if x == count {
			continue
		} else {

		}
	}
}
