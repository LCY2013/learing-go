package string__test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string //字符串默认值是""
	t.Log(s)
	s = "fufeng"
	t.Log(len(s))
	//s[1] = "3" //string是不可变的byte slice
	s = "\xE4\xB8\xA5"
	//s = "\xE4\xBA\xB5\xFF"
	t.Log(s)
	s = "中国"
	t.Log(len(s))
	c := []rune(s)
	t.Logf("中国 unicode %x", c[0])
	t.Logf("中国 utf-8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c,%[1]d", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "a,b,c,d,e"
	parts := strings.Split(s, ",")
	for _, c := range parts {
		t.Log(c)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str " + s)
	if value, ok := strconv.Atoi("10"); ok == nil {
		t.Log(10 + value)
	}
}
