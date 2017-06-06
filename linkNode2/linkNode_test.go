package linkNode2

import (
	// "fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// 创建
func TestCreate(t *testing.T) {
	s := Create(10)
	Convey("create", t, func() {
		Convey("1", func() {
			So(s.Val, ShouldEqual, 0)
			So(s.Next.Val, ShouldEqual, 1)
			So(s.Next.Next.Val, ShouldEqual, 2)
			So(s.Next.Next.Next.Val, ShouldEqual, 3)
		})
	})

	s = s.Shift()
	Convey("删除第一个节点", t, func() {
		Convey("1", func() {
			So(s.Val, ShouldEqual, 1)
			So(s.Next.Val, ShouldEqual, 2)
			So(s.Next.Next.Val, ShouldEqual, 3)
		})
	})

	// s = s.Del(1)
	// Convey("删除指定节点", t, func() {
	// 	Convey("1", func() {
	// 		So(s.Val, ShouldEqual, 1)
	// 		So(s.Next.Val, ShouldEqual, 3)
	// 	})
	// })

}

// 删除第1个节点
