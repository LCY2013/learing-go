package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type funcTimeSpent func(n int) int

func timeSpent(inner funcTimeSpent) funcTimeSpent {
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
