package error

import (
	"errors"
	"fmt"
	"testing"
)

var ParamWrong = errors.New("n should be big then 0")

func Fibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, ParamWrong
	}
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-1]+ret[i-2])
	}
	return ret, nil
}

func TestErrorFibonacci(t *testing.T) {
	if val, err := Fibonacci(-5); err == nil {
		t.Log(val)
	} else {
		t.Log(err)
	}
}

func TestPanic(t *testing.T) {
	//defer func() {
	//	fmt.Println("clear programmer error")
	//}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover error ", err)
		}
	}()

	fmt.Println("start")
	panic(errors.New("occur wrong thing exit"))
	//os.Exit(-1)
}
