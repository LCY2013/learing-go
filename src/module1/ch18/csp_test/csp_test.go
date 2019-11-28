package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() {
	fmt.Println("otherTask running")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("otherTask has finished")
}

func TestSpendTime(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func asyncService() chan string {
	//ch := make(chan string)
	ch := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("return result")
		ch <- ret
		fmt.Println("service executed")
	}()

	return ch
}

func TestAsyncService(t *testing.T) {
	asyncRet := asyncService()
	otherTask()
	fmt.Println(<-asyncRet)
	//fmt.Println(<-asyncService())
	time.Sleep(time.Millisecond * 500)
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 550):
		t.Log("time out")
	}
}
