package type_test

import "testing"

type Int64 int64

func TestImplicit(t *testing.T) {
	var a int32 = 60
	var b int64
	b = int64(a)
	var c Int64
	c = Int64(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPoint := &a
	t.Log(a, aPoint)
	t.Logf("%T %T", a, aPoint)
	t.Log(*aPoint + 1)
}

func TestStr(t *testing.T) {
	var str string
	t.Log("=" + str + "=")
	t.Log(len(str))
	if str == "nil" {
		t.Log("str is nil")
	} else if str == "" {
		t.Log("str is empty")
	}
}
