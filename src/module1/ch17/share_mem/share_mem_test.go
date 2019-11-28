package share_mem

import (
	"sync"
	"testing"
	"time"
)

//协程不安全
func TestShareMemory(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Millisecond * 500)
	t.Log("counter = ", counter)
}

//协程安全
func TestShareMemorySafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Millisecond * 500)
	t.Log("counter = ", counter)
}

func TestShareMemoryWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock()
			counter++
		}()
	}
	wg.Wait()
	t.Log("counter = ", counter)
}
