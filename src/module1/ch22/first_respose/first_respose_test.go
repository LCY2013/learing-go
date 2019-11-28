package first_respose

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//执行任务
func RunTask(n int) string {
	time.Sleep(time.Millisecond * 500)
	return fmt.Sprintf("task runner %d", n)
}

func ExecTask() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(n int) {
			taskId := RunTask(n)
			ch <- taskId
		}(i)
	}
	return <-ch
}

func TestRunTask(t *testing.T) {
	//打印开始前的协程信息
	fmt.Println("before ", runtime.NumGoroutine())
	fmt.Println(ExecTask())
	time.Sleep(time.Millisecond * 500)
	//打印结束的协程信息
	fmt.Println("after ", runtime.NumGoroutine())
}
