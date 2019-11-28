package unsafe

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	var i int64 = 1
	var f float64 = *((*float64)(unsafe.Pointer(&i)))
	t.Log(unsafe.Pointer(&i))
	t.Log(f)
}

type MyInt int64

func TestUnsafeMyInttype(t *testing.T) {
	var i []int64 = []int64{1, 2, 3, 4, 5}
	var m []MyInt = *((*[]MyInt)(unsafe.Pointer(&i)))
	t.Log(unsafe.Pointer(&i))
	t.Log(m)
}

func TestAtomic(t *testing.T) {
	var shareBuffer unsafe.Pointer

	writeDataFn := func() {
		data := []int{0}
		for i := 1; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBuffer, unsafe.Pointer(&data))
	}
	readData := func() {
		data := atomic.LoadPointer(&shareBuffer)
		fmt.Println(data, *(*[]int)(data))
	}

	wg := sync.WaitGroup{}
	//提前开辟地址
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for w := 0; w < 10; w++ {
				writeDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for r := 0; r < 10; r++ {
				readData()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
