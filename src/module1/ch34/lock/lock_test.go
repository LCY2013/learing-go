package lock

import (
	"fmt"
	"sync"
	"testing"
)

var cache map[string]string

const NUMBER_OF_READER int = 40
const READ_TIMES = 100000

func init() {
	cache = make(map[string]string)
	cache["a"] = "a"
	cache["b"] = "b"
}

//不需要锁
func LockFreeAccess() {
	var wg sync.WaitGroup
	wg.Add(NUMBER_OF_READER)

	for i := 0; i < NUMBER_OF_READER; i++ {
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				_, ok := cache["a"]
				if !ok {
					fmt.Println("nothing")
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

//使用读锁
func LockAccess() {
	var wg sync.WaitGroup
	wg.Add(NUMBER_OF_READER)

	//定义读写锁
	rw := new(sync.RWMutex)

	for i := 0; i < NUMBER_OF_READER; i++ {
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				rw.Lock()
				_, ok := cache["a"]
				if !ok {
					fmt.Println("nothing")
				}
				rw.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkLockFreeAccess(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LockFreeAccess()
	}
	b.StopTimer()
}

func BenchmarkLockAccess(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LockAccess()
	}
	b.StopTimer()
}
