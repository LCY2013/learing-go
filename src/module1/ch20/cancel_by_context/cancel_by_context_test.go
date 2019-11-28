package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//是否取消任务
func IsCancelTask(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelTask(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(num int, ctx context.Context) {
			for {
				if IsCancelTask(ctx) {
					break
				}
			}
			fmt.Println("Task", num, "-cancel")
		}(i, ctx)
	}
	time.Sleep(time.Millisecond * 500)
	cancel()
}
