package sync_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 1
		}}
	//第一次获取，没有就会获取默认值
	ret := pool.Get().(int)
	fmt.Println(ret)
	//放入一个值，就不会获取到默认值了
	pool.Put(11)
	runtime.GC() //手动触发一次GC操作，让gc回收sync.pool信息，因为sync.pool和java中的虚引用一样，gc就会回收
	//第二次获取，会获取到11
	ret = pool.Get().(int)
	fmt.Println(ret)
}

func TestGoroutineSyncPool(t *testing.T) {
	//创建sync.pool对象
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 1
		}}
	//先像sync.pool中存放几个变量
	pool.Put(11)
	pool.Put(11)
	pool.Put(11)
	pool.Put(11)
	//pool.Put(11)

	wg := sync.WaitGroup{}
	//启动多个协程去获取池中的数据
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(pool.Get())
			wg.Done()
		}()
	}
	wg.Wait()
}
