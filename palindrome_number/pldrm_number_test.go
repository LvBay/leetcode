package palindrome_number

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetArray(t *testing.T) {
	Convey("get array", t, func() {
		Convey("123", func() {
			So(GetArray(123), ShouldEqual, []int{3, 2, 1})
		})
		Convey("121", func() {
			So(GetArray(121), ShouldEqual, []int{1, 2, 1})
		})
		Convey("-2147483648", func() {
			So(GetArray(-2147483648), ShouldEqual, []int{1, 2, 1})
		})
		Convey("-2147447412", func() {
			So(GetArray(-2147447412), ShouldEqual, []int{1, 2, 1})
		})
	})
}

func TestIsPalindrome(t *testing.T) {
	Convey("isPalindrome", t, func() {
		Convey("111", func() {
			So(isPalindrome(111), ShouldEqual, true)
		})
		Convey("123", func() {
			So(isPalindrome(123), ShouldEqual, false)
		})
		Convey("121", func() {
			So(isPalindrome(121), ShouldEqual, true)
		})
		Convey("-2147483648", func() {
			So(isPalindrome(-2147483648), ShouldEqual, false)
		})
		Convey("-2147447412", func() {
			So(isPalindrome(-2147447412), ShouldEqual, false)
		})
	})
}
