package onece_run

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleton *Singleton
var once sync.Once

//创建单例对象
func CreateSingletonObject() *Singleton {
	once.Do(func() {
		fmt.Println("create singleton object")
		singleton = new(Singleton)
	})
	return singleton
}

func TestCreateSingletonObject(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			singletonObject := CreateSingletonObject()
			fmt.Printf("%x \n", unsafe.Pointer(singletonObject))
			//fmt.Println(&singletonObject)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
