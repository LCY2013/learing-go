package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

//生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//关闭chan
		close(ch)
		wg.Done()
	}()
}

//消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println("ch -> ", data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func dataReceiver1(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println("ch -> ", data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver1(ch, &wg)
	wg.Wait()
}
