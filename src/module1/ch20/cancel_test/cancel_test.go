package cancel_test

import (
	"fmt"
	"testing"
	"time"
)

//是否取消任务
func IsCancelTask(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

//取消其中某一个任务
func Cancel_1(ch chan struct{}) {
	ch <- struct{}{}
}

//关闭通道，通知所以的任务
func Cancel_2(ch chan struct{}) {
	close(ch)
}

func TestCancelTask(t *testing.T) {
	ch := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(num int, ch chan struct{}) {
			for {
				if IsCancelTask(ch) {
					break
				}
			}
			fmt.Println("Task", num, "-cancel")
		}(i, ch)
	}
	//Cancel_1(ch)
	Cancel_2(ch)
	time.Sleep(time.Millisecond * 500)
}
