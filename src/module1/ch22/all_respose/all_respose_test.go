package all_respose

import (
	"fmt"
	"runtime"
	"testing"
)

//执行任务
func RunTask(n int) string {
	//time.Sleep(time.Millisecond * 500)
	return fmt.Sprintf("task runner %d\n", n)
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
	result := ""
	for i := 0; i < numOfRunner; i++ {
		result += <-ch
	}
	return result
}

func TestRunTask(t *testing.T) {
	//打印开始前的协程信息
	fmt.Println("before ", runtime.NumGoroutine())
	fmt.Println(ExecTask())
	//打印结束的协程信息
	fmt.Println("after ", runtime.NumGoroutine())
}
