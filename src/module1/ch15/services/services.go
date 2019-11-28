package services

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("init3")
}

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

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
