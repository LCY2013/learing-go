package unit

import (
	"bytes"
	"testing"
)

//单元测试
//func TestSquare(t *testing.T) {
//	input := [...]int{1, 2, 3, 4, 5}
//	expect := [...]int{1, 4, 9, 169, 25}
//
//	for i := 0; i < len(input); i++ {
//		ret := Square(input[i])
//		if expect[i] != ret {
//			t.Errorf("the input is %d,the expect is %d,the actual is %d",
//				input[i], expect[i], ret)
//		}
//	}
//}
//
//func TestErrorInCode(t *testing.T) {
//	fmt.Println("start")
//	t.Error("error")
//	fmt.Println("end")
//}
//
//func TestFatalInCode(t *testing.T) {
//	fmt.Println("start")
//	t.Fatal("fatal")
//	fmt.Println("end")
//}
//
////利用第三方库做断言
//func TestSquareWithAssert(t *testing.T) {
//	input := [...]int{1, 2, 3, 4, 5, 6}
//	expect := [...]int{1, 4, 9, 16, 125, 36}
//	for i := 0; i < len(input); i++ {
//		ret := Square(input[i])
//		assert.EqualValues(t, expect[i], ret)
//	}
//}
//
//func TestConcatStringByAdd(t *testing.T) {
//	input := [...]string{"1", "2", "3", "4", "5"}
//	ret := ""
//	for i := 0; i < len(input); i++ {
//		ret += input[i]
//	}
//	assert.Equal(t, "12345", ret)
//}
//
//func TestConcatStringByBytesBuffer(t *testing.T) {
//	asserts := assert.New(t)
//	input := [...]string{"1", "2", "3", "4", "5"}
//	var buf bytes.Buffer
//	for i := 0; i < len(input); i++ {
//		buf.WriteString(input[i])
//	}
//	asserts.Equal("12345", buf.String())
//}

//benchmark 测试
func BenchmarkConcatStringByAdd(b *testing.B) {
	input := [...]string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, val := range input {
			ret += val
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBuffer(b *testing.B) {
	input := [...]string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, val := range input {
			buf.WriteString(val)
		}
	}
	b.StopTimer()
}
