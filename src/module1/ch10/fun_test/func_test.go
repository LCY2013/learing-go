package fun_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(100)
}

func TestFn(t *testing.T) {
	key1, key2 := returnMultiValue()
	t.Log(key1, key2)
}

func timeSpent(inner func(n int) int) func(n int) int {
	return func(n int) int {
		now := time.Now()
		result := inner(n)
		fmt.Println("time spent:", time.Since(now).Seconds())
		return result
	}
}

func slowFunc(n int) int {
	time.Sleep(time.Second * 2)
	return n
}

func TestTimeSpent(t *testing.T) {
	spentFunction := timeSpent(slowFunc)
	t.Log(spentFunction(1995))
}

func sum(ops ...int) int {
	ret := 0
	for _, value := range ops {
		ret += value
	}
	return ret
}

//可变参数
func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3))
	t.Log(sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

//defer 延迟执行函数
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start")
	//panic("err")
}
