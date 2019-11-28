package bdd

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//在本代码目录执行 /Users/a/softwork/go/gopath/bin/goconvey  可以浏览器查看结果信息
// $GOPATH/bin/goconvey
func TestSpec(t *testing.T) {

	//only pass t into top-level Convey Calls
	Convey("Given 2 even numbers ", t, func() {
		n1 := 1
		n2 := 3
		Convey("When add the two numbers", func() {
			n3 := n1 + n2
			Convey("Then the result is still even", func() {
				So(n3%2, ShouldEqual, 1)
			})
		})
	})
}
