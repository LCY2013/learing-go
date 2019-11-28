package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}
	e := [...]int{1, 3, 4, 2}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
	t.Log(a == e)
}

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 //0111
	// &^ 按位清零，如果右边的数据在位置上是0，左边就保留，右边是1，左边就置为0
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
